// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

/*
Package delay provides a way to execute code outside the scope of a
user request by using the taskqueue API.

To declare a function that may be executed later, call Func
in a top-level assignment context, passing it an arbitrary string key
and a function whose first argument is of type appengine.Context.
	var laterFunc = delay.Func("key", myFunc)
It is also possible to use a function literal.
	var laterFunc = delay.Func("key", func(c appengine.Context, x string) {
		// ...
	})

To call a function, invoke its Call method.
	laterFunc.Call(c, "something")
A function may be called any number of times.

The arguments to functions may be of any type that is encodable by the
gob package.

Any errors during initialization or execution of a function will be
logged to the application logs. Error logs that occur during initialization will
be associated with the request that invoked the Call method.

The state of a function invocation that has not yet successfully
executed is preserved by combining the file name in which it is declared
with the string key that was passed to the Func function. Updating an app
with pending function invocations is safe as long as the relevant
functions have the (filename, key) combination preserved.
*/
package delay

import (
	"bytes"
	"gob"
	"http"
	"os"
	"reflect"
	"runtime"

	"appengine"
	"appengine/taskqueue"
)

// Function represents a function that may have a delayed invocation.
type Function struct {
	fv  reflect.Value // Kind() == reflect.Func
	key string
	err os.Error // any error during initialization
}

const (
	// The HTTP path for invocations.
	path = "/_ah/queue/go/delay"
	// Use the default queue.
	queue = ""
)

var (
	// registry of all delayed functions
	funcs = make(map[string]*Function)

	// precomputed type
	contextType = reflect.TypeOf((*appengine.Context)(nil)).Elem()
)

// Func declares a new Function. The second argument must be a function with a
// first argument of type appengine.Context.
// This function must be called at program initialization time. That means it
// must be called in a global variable declaration or from an init function.
// This restriction is necessary because the instance that delays a function
// call may not be the one that executes it. Only the code executed at program
// initialization time is guaranteed to have been run by an instance before it
// receives a request.
func Func(key string, i interface{}) *Function {
	f := &Function{fv: reflect.ValueOf(i)}

	// Derive unique, somewhat stable key for this func.
	_, file, _, _ := runtime.Caller(1)
	f.key = file + ":" + key

	t := f.fv.Type()
	if t.Kind() != reflect.Func {
		f.err = os.ErrorString("not a function")
		return f
	}
	if t.NumIn() == 0 || t.In(0) != contextType {
		f.err = os.ErrorString("first argument must be appengine.Context")
		return f
	}

	funcs[f.key] = f
	return f
}

type invocation struct {
	Key  string
	Args []interface{}
}

// Call invokes a delayed function.
// Note that the function will be executed later.
func (f *Function) Call(c appengine.Context, args ...interface{}) {
	if f.err != nil {
		c.Errorf("delay: func is invalid: %v", f.err)
		return
	}

	nArgs := len(args) + 1 // +1 for the appengine.Context
	ft := f.fv.Type()
	minArgs := ft.NumIn()
	if ft.IsVariadic() {
		minArgs--
	}
	if nArgs < minArgs {
		c.Errorf("delay: too few arguments to func: %d < %d", nArgs, minArgs)
		return
	}
	if !ft.IsVariadic() && nArgs > minArgs {
		c.Errorf("delay: too many arguments to func: %d > %d", nArgs, minArgs)
		return
	}

	// TODO: check arg types.

	inv := invocation{
		Key:  f.key,
		Args: args,
	}

	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(inv); err != nil {
		c.Errorf("delay: gob encoding failed: %v", err)
		return
	}

	task := &taskqueue.Task{
		Path:    path,
		Payload: buf.Bytes(),
	}
	if _, err := taskqueue.Add(c, task, queue); err != nil {
		c.Errorf("delay: taskqueue.Add failed: %v", err)
		return
	}
}

func init() {
	http.HandleFunc(path, runFunc)
}

func runFunc(w http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	defer req.Body.Close()

	var inv invocation
	if err := gob.NewDecoder(req.Body).Decode(&inv); err != nil {
		c.Errorf("delay: failed decoding task payload: %v", err)
		c.Warningf("delay: dropping task")
		return
	}

	f := funcs[inv.Key]
	if f == nil {
		c.Errorf("delay: no func with key %q found", inv.Key)
		c.Warningf("delay: dropping task")
		return
	}

	// TODO: This is broken for variadic functions.
	in := make([]reflect.Value, f.fv.Type().NumIn())
	in[0] = reflect.ValueOf(c)
	for i := 1; i < len(in); i++ {
		in[i] = reflect.ValueOf(inv.Args[i-1])
	}
	f.fv.Call(in)
}
