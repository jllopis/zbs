// Code generated by protoc-gen-go.
// source: services/common.proto
// DO NOT EDIT!

/*
Package services is a generated protocol buffer package.

It is generated from these files:
	services/common.proto
	services/time.proto
	services/zbs.proto

It has these top-level messages:
	IdMessage
	UuidMessage
	Timestamp
	ServerTimeMessage
	JobArrayMsg
	JobFilterMsg
	JobSortingMsg
	JobMsg
	JobDataMsg
	ScriptMsg
*/
package services

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IdMessage struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *IdMessage) Reset()                    { *m = IdMessage{} }
func (m *IdMessage) String() string            { return proto.CompactTextString(m) }
func (*IdMessage) ProtoMessage()               {}
func (*IdMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IdMessage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UuidMessage struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *UuidMessage) Reset()                    { *m = UuidMessage{} }
func (m *UuidMessage) String() string            { return proto.CompactTextString(m) }
func (*UuidMessage) ProtoMessage()               {}
func (*UuidMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UuidMessage) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type Timestamp struct {
	// Representa el tiempo Unix (desde epoch) en UTC
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func init() {
	proto.RegisterType((*IdMessage)(nil), "services.IdMessage")
	proto.RegisterType((*UuidMessage)(nil), "services.UuidMessage")
	proto.RegisterType((*Timestamp)(nil), "services.Timestamp")
}

func init() { proto.RegisterFile("services/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x80, 0x09, 0x4b, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16,
	0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0xd4, 0x29,
	0x49, 0x73, 0x71, 0x7a, 0xa6, 0xf8, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0xf1, 0x71, 0x31,
	0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x31, 0x65, 0xa6, 0x28, 0x29, 0x72, 0x71,
	0x87, 0x96, 0x66, 0xc2, 0xa5, 0x85, 0xb8, 0x58, 0x4a, 0x4b, 0xa1, 0x0a, 0x38, 0x83, 0xc0, 0x6c,
	0x25, 0x55, 0x2e, 0xce, 0x90, 0xcc, 0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x21, 0x09, 0x2e,
	0xf6, 0xe2, 0xd4, 0xe4, 0xfc, 0xbc, 0x94, 0x62, 0xa8, 0x21, 0x30, 0x6e, 0x12, 0x1b, 0xd8, 0x36,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xa5, 0x7d, 0x55, 0xae, 0x00, 0x00, 0x00,
}
