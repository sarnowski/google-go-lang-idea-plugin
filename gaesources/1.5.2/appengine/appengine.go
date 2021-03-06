// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package appengine provides functionality that is common across
// App Engine APIs.
package appengine

import (
	"http"
	"os"

	"appengine_internal"
)

// IsDevAppServer reports whether the App Engine app is running in the
// development App Server.
func IsDevAppServer() bool {
	return appengine_internal.IsDevAppServer()
}

// Context represents the context of an in-flight HTTP request.
type Context interface {
	// Debugf formats its arguments according to the format, analogous to fmt.Printf,
	// and records the text as a log message at Debug level.
	Debugf(format string, args ...interface{})

	// Infof is like Debugf, but at Info level.
	Infof(format string, args ...interface{})

	// Warningf is like Debugf, but at Warning level.
	Warningf(format string, args ...interface{})

	// Errorf is like Debugf, but at Error level.
	Errorf(format string, args ...interface{})

	// Criticalf is like Debugf, but at Critical level.
	Criticalf(format string, args ...interface{})

	// AppID returns the application ID for the current application.
	// The string will be a plain application ID (e.g. "appid"),
	// with a domain prefix for custom domain deployments (e.g. "example.com:appid").
	AppID() string

	// The remaining methods are for internal use only.
	// Developer-facing APIs wrap these methods to provide a more friendly API.

	Call(service, method string, in, out interface{}) os.Error
	FullyQualifiedAppID() string
	Request() interface{}
}

// NewContext returns a new context for an in-flight HTTP request.
func NewContext(req *http.Request) Context {
	return appengine_internal.NewContext(req)
}

// BlobKey is a key for a blobstore blob.
//
// Conceptually, this type belongs in the blobstore package, but it lives in
// the appengine package to avoid a circular dependency: blobstore depends on
// datastore, and datastore needs to refer to the BlobKey type.
type BlobKey string
