// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/examples/template/srv/proto/example/example.proto

/*
Package go_micro_srv_template is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/examples/template/srv/proto/example/example.proto

It has these top-level messages:
	Message
	Request
	Response
	StreamingRequest
	StreamingResponse
	Ping
	Pong
*/
package go_micro_srv_template

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Say string `protobuf:"bytes,1,opt,name=say" json:"say,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *StreamingRequest) Reset()                    { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()               {}
func (*StreamingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *StreamingResponse) Reset()                    { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()               {}
func (*StreamingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke int64 `protobuf:"varint,1,opt,name=stroke" json:"stroke,omitempty"`
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke int64 `protobuf:"varint,1,opt,name=stroke" json:"stroke,omitempty"`
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.template.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.template.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.template.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.template.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.template.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.template.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.template.Pong")
}

func init() {
	proto.RegisterFile("github.com/micro/examples/template/srv/proto/example/example.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x1b, 0x5b, 0xdb, 0x3a, 0xa7, 0xba, 0xa8, 0x48, 0xab, 0x55, 0xf6, 0x62, 0xbc, 0x6c,
	0x8a, 0xbe, 0x81, 0xe2, 0x41, 0x41, 0x90, 0xf8, 0x00, 0xb2, 0x0d, 0xc3, 0x1a, 0xcc, 0xee, 0xc6,
	0x9d, 0x4d, 0xd1, 0xa3, 0x6f, 0x2e, 0xd9, 0x24, 0x22, 0xd2, 0xe0, 0x29, 0x33, 0xf9, 0xfe, 0x7f,
	0x98, 0xf9, 0x59, 0xb8, 0x51, 0xb9, 0x7f, 0xad, 0xd6, 0x22, 0xb3, 0x3a, 0xd1, 0x79, 0xe6, 0x6c,
	0x82, 0x1f, 0x52, 0x97, 0x05, 0x52, 0xe2, 0x51, 0x97, 0x85, 0xf4, 0x98, 0x90, 0xdb, 0x24, 0xa5,
	0xb3, 0xfe, 0x87, 0x75, 0x5f, 0x11, 0xfe, 0xb2, 0x43, 0x65, 0x45, 0xf0, 0x0a, 0x72, 0x1b, 0xd1,
	0xd9, 0xf8, 0x02, 0x26, 0x8f, 0x48, 0x24, 0x15, 0xb2, 0x19, 0x0c, 0x49, 0x7e, 0x1e, 0x47, 0xe7,
	0x51, 0xbc, 0x97, 0xd6, 0x25, 0x3f, 0x85, 0x49, 0x8a, 0xef, 0x15, 0x92, 0x67, 0x0c, 0x46, 0x46,
	0x6a, 0x6c, 0x69, 0xa8, 0xf9, 0x09, 0x4c, 0x53, 0xa4, 0xd2, 0x1a, 0x0a, 0x66, 0x4d, 0xaa, 0x33,
	0x6b, 0x52, 0x3c, 0x86, 0xd9, 0xb3, 0x77, 0x28, 0x75, 0x6e, 0x54, 0x37, 0xe5, 0x00, 0x76, 0x33,
	0x5b, 0x19, 0x1f, 0x74, 0xc3, 0xb4, 0x69, 0xf8, 0x25, 0xec, 0xff, 0x52, 0xb6, 0x03, 0xb7, 0x4b,
	0x97, 0x30, 0x7a, 0xca, 0x8d, 0x62, 0x47, 0x30, 0x26, 0xef, 0xec, 0x1b, 0xb6, 0xb8, 0xed, 0x02,
	0xb7, 0xfd, 0xfc, 0xea, 0x6b, 0x07, 0x26, 0x77, 0x4d, 0x2e, 0xec, 0x1e, 0x46, 0xb7, 0xb2, 0x28,
	0xd8, 0x52, 0x6c, 0x8d, 0x46, 0xb4, 0x4b, 0xcf, 0xcf, 0x7a, 0x79, 0xb3, 0x2a, 0x1f, 0xb0, 0x17,
	0x18, 0x37, 0x17, 0xb0, 0x8b, 0x1e, 0xf1, 0xdf, 0x28, 0xe6, 0xf1, 0xff, 0xc2, 0x6e, 0xfc, 0x2a,
	0x62, 0x0f, 0x30, 0xad, 0xef, 0x0e, 0xb7, 0x2d, 0x7a, 0x9c, 0xb5, 0x60, 0xde, 0x0b, 0xad, 0x51,
	0x7c, 0x10, 0x47, 0xab, 0x68, 0x3d, 0x0e, 0x0f, 0xe2, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x97,
	0xcb, 0x75, 0xaa, 0x56, 0x02, 0x00, 0x00,
}
