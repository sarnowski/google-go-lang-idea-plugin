// Code generated by protoc-gen-go from "capability_service.proto"
// DO NOT EDIT!

package appengine

import proto "goprotobuf.googlecode.com/hg/proto"
import "math"
import "os"


// Reference proto, math & os imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf
var _ os.Error


type CapabilityConfig_Status int32

const (
	CapabilityConfig_ENABLED   = 1
	CapabilityConfig_SCHEDULED = 2
	CapabilityConfig_DISABLED  = 3
	CapabilityConfig_UNKNOWN   = 4
)

var CapabilityConfig_Status_name = map[int32]string{
	1: "ENABLED",
	2: "SCHEDULED",
	3: "DISABLED",
	4: "UNKNOWN",
}
var CapabilityConfig_Status_value = map[string]int32{
	"ENABLED":   1,
	"SCHEDULED": 2,
	"DISABLED":  3,
	"UNKNOWN":   4,
}

func NewCapabilityConfig_Status(x int32) *CapabilityConfig_Status {
	e := CapabilityConfig_Status(x)
	return &e
}
func (x CapabilityConfig_Status) String() string {
	return proto.EnumName(CapabilityConfig_Status_name, int32(x))
}

type IsEnabledResponse_SummaryStatus int32

const (
	IsEnabledResponse_ENABLED          = 1
	IsEnabledResponse_SCHEDULED_FUTURE = 2
	IsEnabledResponse_SCHEDULED_NOW    = 3
	IsEnabledResponse_DISABLED         = 4
	IsEnabledResponse_UNKNOWN          = 5
)

var IsEnabledResponse_SummaryStatus_name = map[int32]string{
	1: "ENABLED",
	2: "SCHEDULED_FUTURE",
	3: "SCHEDULED_NOW",
	4: "DISABLED",
	5: "UNKNOWN",
}
var IsEnabledResponse_SummaryStatus_value = map[string]int32{
	"ENABLED":          1,
	"SCHEDULED_FUTURE": 2,
	"SCHEDULED_NOW":    3,
	"DISABLED":         4,
	"UNKNOWN":          5,
}

func NewIsEnabledResponse_SummaryStatus(x int32) *IsEnabledResponse_SummaryStatus {
	e := IsEnabledResponse_SummaryStatus(x)
	return &e
}
func (x IsEnabledResponse_SummaryStatus) String() string {
	return proto.EnumName(IsEnabledResponse_SummaryStatus_name, int32(x))
}

type CapabilityConfigList struct {
	Config           []*CapabilityConfig "PB(bytes,1,rep,name=config)"
	DefaultConfig    *CapabilityConfig   "PB(bytes,2,opt,name=default_config)"
	XXX_unrecognized []byte
}

func (this *CapabilityConfigList) Reset()         { *this = CapabilityConfigList{} }
func (this *CapabilityConfigList) String() string { return proto.CompactTextString(this) }

type CapabilityConfig struct {
	Package          *string                  "PB(bytes,1,req,name=package)"
	Capability       *string                  "PB(bytes,2,req,name=capability)"
	Status           *CapabilityConfig_Status "PB(varint,3,opt,name=status,enum=appengine.CapabilityConfig_Status,def=4)"
	ScheduledTime    *string                  "PB(bytes,7,opt,name=scheduled_time)"
	InternalMessage  *string                  "PB(bytes,4,opt,name=internal_message)"
	AdminMessage     *string                  "PB(bytes,5,opt,name=admin_message)"
	ErrorMessage     *string                  "PB(bytes,6,opt,name=error_message)"
	XXX_unrecognized []byte
}

func (this *CapabilityConfig) Reset()         { *this = CapabilityConfig{} }
func (this *CapabilityConfig) String() string { return proto.CompactTextString(this) }

const Default_CapabilityConfig_Status CapabilityConfig_Status = CapabilityConfig_UNKNOWN

type IsEnabledRequest struct {
	Package          *string  "PB(bytes,1,req,name=package)"
	Capability       []string "PB(bytes,2,rep,name=capability)"
	Call             []string "PB(bytes,3,rep,name=call)"
	XXX_unrecognized []byte
}

func (this *IsEnabledRequest) Reset()         { *this = IsEnabledRequest{} }
func (this *IsEnabledRequest) String() string { return proto.CompactTextString(this) }

type IsEnabledResponse struct {
	SummaryStatus      *IsEnabledResponse_SummaryStatus "PB(varint,1,req,name=summary_status,enum=appengine.IsEnabledResponse_SummaryStatus)"
	TimeUntilScheduled *int64                           "PB(varint,2,opt,name=time_until_scheduled)"
	Config             []*CapabilityConfig              "PB(bytes,3,rep,name=config)"
	XXX_unrecognized   []byte
}

func (this *IsEnabledResponse) Reset()         { *this = IsEnabledResponse{} }
func (this *IsEnabledResponse) String() string { return proto.CompactTextString(this) }

func init() {
	proto.RegisterEnum("appengine.CapabilityConfig_Status", CapabilityConfig_Status_name, CapabilityConfig_Status_value)
	proto.RegisterEnum("appengine.IsEnabledResponse_SummaryStatus", IsEnabledResponse_SummaryStatus_name, IsEnabledResponse_SummaryStatus_value)
}