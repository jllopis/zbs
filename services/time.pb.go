// Code generated by protoc-gen-go.
// source: services/time.proto
// DO NOT EDIT!

package services

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ServerTimeMessage struct {
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *ServerTimeMessage) Reset()                    { *m = ServerTimeMessage{} }
func (m *ServerTimeMessage) String() string            { return proto.CompactTextString(m) }
func (*ServerTimeMessage) ProtoMessage()               {}
func (*ServerTimeMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ServerTimeMessage) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*ServerTimeMessage)(nil), "services.ServerTimeMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TimeService service

type TimeServiceClient interface {
	GetServerTime(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ServerTimeMessage, error)
}

type timeServiceClient struct {
	cc *grpc.ClientConn
}

func NewTimeServiceClient(cc *grpc.ClientConn) TimeServiceClient {
	return &timeServiceClient{cc}
}

func (c *timeServiceClient) GetServerTime(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ServerTimeMessage, error) {
	out := new(ServerTimeMessage)
	err := grpc.Invoke(ctx, "/services.TimeService/GetServerTime", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TimeService service

type TimeServiceServer interface {
	GetServerTime(context.Context, *google_protobuf1.Empty) (*ServerTimeMessage, error)
}

func RegisterTimeServiceServer(s *grpc.Server, srv TimeServiceServer) {
	s.RegisterService(&_TimeService_serviceDesc, srv)
}

func _TimeService_GetServerTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).GetServerTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.TimeService/GetServerTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).GetServerTime(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TimeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.TimeService",
	HandlerType: (*TimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerTime",
			Handler:    _TimeService_GetServerTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/time.proto",
}

func init() { proto.RegisterFile("services/time.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2f, 0xc9, 0xcc, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x80, 0x09, 0x4a, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea,
	0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0xd4, 0x49, 0x49, 0x43,
	0x65, 0xc1, 0xbc, 0xa4, 0xd2, 0x34, 0xfd, 0xd4, 0xdc, 0x82, 0x92, 0x4a, 0x88, 0xa4, 0x92, 0x26,
	0x97, 0x60, 0x70, 0x6a, 0x51, 0x59, 0x6a, 0x51, 0x48, 0x66, 0x6e, 0xaa, 0x6f, 0x6a, 0x71, 0x71,
	0x62, 0x7a, 0xaa, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x73, 0x10, 0x84, 0x63, 0x94, 0xca, 0xc5, 0x0d, 0x52, 0x14, 0x0c, 0xb1, 0x55, 0x28, 0x8c,
	0x8b, 0xd7, 0x3d, 0xb5, 0x04, 0xa1, 0x59, 0x48, 0x4c, 0x0f, 0x62, 0x91, 0x1e, 0xcc, 0x22, 0x3d,
	0x57, 0x90, 0x45, 0x52, 0xd2, 0x7a, 0x30, 0x87, 0xea, 0x61, 0x58, 0xa5, 0x24, 0xd0, 0x74, 0xf9,
	0xc9, 0x64, 0x26, 0x2e, 0x21, 0x0e, 0xfd, 0x32, 0x43, 0xb0, 0xe7, 0x92, 0xd8, 0xc0, 0xda, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x14, 0xd7, 0xaa, 0x57, 0xf4, 0x00, 0x00, 0x00,
}
