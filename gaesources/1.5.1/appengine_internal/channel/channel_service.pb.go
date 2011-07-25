// Code generated by protoc-gen-go from "channel_service.proto"
// DO NOT EDIT!

package appengine

import proto "goprotobuf.googlecode.com/hg/proto"
import "math"
import "os"


// Reference proto, math & os imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf
var _ os.Error


type ChannelServiceError_ErrorCode int32

const (
	ChannelServiceError_OK                  = 0
	ChannelServiceError_INTERNAL_ERROR      = 1
	ChannelServiceError_INVALID_CHANNEL_KEY = 2
	ChannelServiceError_BAD_MESSAGE         = 3
)

var ChannelServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
	2: "INVALID_CHANNEL_KEY",
	3: "BAD_MESSAGE",
}
var ChannelServiceError_ErrorCode_value = map[string]int32{
	"OK":                  0,
	"INTERNAL_ERROR":      1,
	"INVALID_CHANNEL_KEY": 2,
	"BAD_MESSAGE":         3,
}

func NewChannelServiceError_ErrorCode(x int32) *ChannelServiceError_ErrorCode {
	e := ChannelServiceError_ErrorCode(x)
	return &e
}
func (x ChannelServiceError_ErrorCode) String() string {
	return proto.EnumName(ChannelServiceError_ErrorCode_name, int32(x))
}

type ChannelServiceError struct {
	XXX_unrecognized []byte
}

func (this *ChannelServiceError) Reset()         { *this = ChannelServiceError{} }
func (this *ChannelServiceError) String() string { return proto.CompactTextString(this) }

type CreateChannelRequest struct {
	ApplicationKey   *string "PB(bytes,1,req,name=application_key)"
	XXX_unrecognized []byte
}

func (this *CreateChannelRequest) Reset()         { *this = CreateChannelRequest{} }
func (this *CreateChannelRequest) String() string { return proto.CompactTextString(this) }

type CreateChannelResponse struct {
	ClientId         *string "PB(bytes,2,opt,name=client_id)"
	XXX_unrecognized []byte
}

func (this *CreateChannelResponse) Reset()         { *this = CreateChannelResponse{} }
func (this *CreateChannelResponse) String() string { return proto.CompactTextString(this) }

type SendMessageRequest struct {
	ApplicationKey   *string "PB(bytes,1,req,name=application_key)"
	Message          *string "PB(bytes,2,req,name=message)"
	XXX_unrecognized []byte
}

func (this *SendMessageRequest) Reset()         { *this = SendMessageRequest{} }
func (this *SendMessageRequest) String() string { return proto.CompactTextString(this) }

func init() {
	proto.RegisterEnum("appengine.ChannelServiceError_ErrorCode", ChannelServiceError_ErrorCode_name, ChannelServiceError_ErrorCode_value)
}
