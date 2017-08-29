// Code generated by protoc-gen-go. DO NOT EDIT.
// source: desc_test_proto3.proto

package testprotos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import jhump_protoreflect_desc "github.com/jhump/protoreflect/internal/testprotos/pkg"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Proto3Enum int32

const (
	Proto3Enum_UNKNOWN Proto3Enum = 0
	Proto3Enum_VALUE1  Proto3Enum = 1
	Proto3Enum_VALUE2  Proto3Enum = 2
)

var Proto3Enum_name = map[int32]string{
	0: "UNKNOWN",
	1: "VALUE1",
	2: "VALUE2",
}
var Proto3Enum_value = map[string]int32{
	"UNKNOWN": 0,
	"VALUE1":  1,
	"VALUE2":  2,
}

func (x Proto3Enum) String() string {
	return proto.EnumName(Proto3Enum_name, int32(x))
}
func (Proto3Enum) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type TestRequest struct {
	Foo   []Proto3Enum                                    `protobuf:"varint,1,rep,packed,name=foo,enum=testprotos.Proto3Enum" json:"foo,omitempty"`
	Bar   string                                          `protobuf:"bytes,2,opt,name=bar" json:"bar,omitempty"`
	Baz   *TestMessage                                    `protobuf:"bytes,3,opt,name=baz" json:"baz,omitempty"`
	Snafu *TestMessage_NestedMessage_AnotherNestedMessage `protobuf:"bytes,4,opt,name=snafu" json:"snafu,omitempty"`
}

func (m *TestRequest) Reset()                    { *m = TestRequest{} }
func (m *TestRequest) String() string            { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()               {}
func (*TestRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *TestRequest) GetFoo() []Proto3Enum {
	if m != nil {
		return m.Foo
	}
	return nil
}

func (m *TestRequest) GetBar() string {
	if m != nil {
		return m.Bar
	}
	return ""
}

func (m *TestRequest) GetBaz() *TestMessage {
	if m != nil {
		return m.Baz
	}
	return nil
}

func (m *TestRequest) GetSnafu() *TestMessage_NestedMessage_AnotherNestedMessage {
	if m != nil {
		return m.Snafu
	}
	return nil
}

type TestResponse struct {
	Atm *AnotherTestMessage `protobuf:"bytes,1,opt,name=atm" json:"atm,omitempty"`
}

func (m *TestResponse) Reset()                    { *m = TestResponse{} }
func (m *TestResponse) String() string            { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()               {}
func (*TestResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *TestResponse) GetAtm() *AnotherTestMessage {
	if m != nil {
		return m.Atm
	}
	return nil
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "testprotos.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "testprotos.TestResponse")
	proto.RegisterEnum("testprotos.Proto3Enum", Proto3Enum_name, Proto3Enum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	DoSomething(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*jhump_protoreflect_desc.Bar, error)
	DoSomethingElse(ctx context.Context, opts ...grpc.CallOption) (TestService_DoSomethingElseClient, error)
	DoSomethingAgain(ctx context.Context, in *jhump_protoreflect_desc.Bar, opts ...grpc.CallOption) (TestService_DoSomethingAgainClient, error)
	DoSomethingForever(ctx context.Context, opts ...grpc.CallOption) (TestService_DoSomethingForeverClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) DoSomething(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*jhump_protoreflect_desc.Bar, error) {
	out := new(jhump_protoreflect_desc.Bar)
	err := grpc.Invoke(ctx, "/testprotos.TestService/DoSomething", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) DoSomethingElse(ctx context.Context, opts ...grpc.CallOption) (TestService_DoSomethingElseClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/testprotos.TestService/DoSomethingElse", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceDoSomethingElseClient{stream}
	return x, nil
}

type TestService_DoSomethingElseClient interface {
	Send(*TestMessage) error
	CloseAndRecv() (*TestResponse, error)
	grpc.ClientStream
}

type testServiceDoSomethingElseClient struct {
	grpc.ClientStream
}

func (x *testServiceDoSomethingElseClient) Send(m *TestMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceDoSomethingElseClient) CloseAndRecv() (*TestResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) DoSomethingAgain(ctx context.Context, in *jhump_protoreflect_desc.Bar, opts ...grpc.CallOption) (TestService_DoSomethingAgainClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[1], c.cc, "/testprotos.TestService/DoSomethingAgain", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceDoSomethingAgainClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_DoSomethingAgainClient interface {
	Recv() (*AnotherTestMessage, error)
	grpc.ClientStream
}

type testServiceDoSomethingAgainClient struct {
	grpc.ClientStream
}

func (x *testServiceDoSomethingAgainClient) Recv() (*AnotherTestMessage, error) {
	m := new(AnotherTestMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) DoSomethingForever(ctx context.Context, opts ...grpc.CallOption) (TestService_DoSomethingForeverClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[2], c.cc, "/testprotos.TestService/DoSomethingForever", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceDoSomethingForeverClient{stream}
	return x, nil
}

type TestService_DoSomethingForeverClient interface {
	Send(*TestRequest) error
	Recv() (*TestResponse, error)
	grpc.ClientStream
}

type testServiceDoSomethingForeverClient struct {
	grpc.ClientStream
}

func (x *testServiceDoSomethingForeverClient) Send(m *TestRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceDoSomethingForeverClient) Recv() (*TestResponse, error) {
	m := new(TestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	DoSomething(context.Context, *TestRequest) (*jhump_protoreflect_desc.Bar, error)
	DoSomethingElse(TestService_DoSomethingElseServer) error
	DoSomethingAgain(*jhump_protoreflect_desc.Bar, TestService_DoSomethingAgainServer) error
	DoSomethingForever(TestService_DoSomethingForeverServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_DoSomething_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).DoSomething(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testprotos.TestService/DoSomething",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).DoSomething(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_DoSomethingElse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).DoSomethingElse(&testServiceDoSomethingElseServer{stream})
}

type TestService_DoSomethingElseServer interface {
	SendAndClose(*TestResponse) error
	Recv() (*TestMessage, error)
	grpc.ServerStream
}

type testServiceDoSomethingElseServer struct {
	grpc.ServerStream
}

func (x *testServiceDoSomethingElseServer) SendAndClose(m *TestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceDoSomethingElseServer) Recv() (*TestMessage, error) {
	m := new(TestMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_DoSomethingAgain_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(jhump_protoreflect_desc.Bar)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).DoSomethingAgain(m, &testServiceDoSomethingAgainServer{stream})
}

type TestService_DoSomethingAgainServer interface {
	Send(*AnotherTestMessage) error
	grpc.ServerStream
}

type testServiceDoSomethingAgainServer struct {
	grpc.ServerStream
}

func (x *testServiceDoSomethingAgainServer) Send(m *AnotherTestMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_DoSomethingForever_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).DoSomethingForever(&testServiceDoSomethingForeverServer{stream})
}

type TestService_DoSomethingForeverServer interface {
	Send(*TestResponse) error
	Recv() (*TestRequest, error)
	grpc.ServerStream
}

type testServiceDoSomethingForeverServer struct {
	grpc.ServerStream
}

func (x *testServiceDoSomethingForeverServer) Send(m *TestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceDoSomethingForeverServer) Recv() (*TestRequest, error) {
	m := new(TestRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testprotos.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSomething",
			Handler:    _TestService_DoSomething_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoSomethingElse",
			Handler:       _TestService_DoSomethingElse_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DoSomethingAgain",
			Handler:       _TestService_DoSomethingAgain_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DoSomethingForever",
			Handler:       _TestService_DoSomethingForever_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "desc_test_proto3.proto",
}

func init() { proto.RegisterFile("desc_test_proto3.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xd9, 0x18, 0x8a, 0x98, 0x20, 0xb0, 0xe6, 0xd0, 0x5a, 0x16, 0x42, 0x51, 0x4f, 0x86,
	0xc3, 0x3a, 0x71, 0x6f, 0x9c, 0x48, 0xd5, 0xf4, 0x52, 0x30, 0x95, 0x4b, 0x41, 0xe2, 0x52, 0x6d,
	0xdc, 0x89, 0x6d, 0x1a, 0xef, 0x9a, 0xdd, 0x75, 0x0f, 0x7d, 0x36, 0xce, 0x3c, 0x17, 0xb2, 0xb7,
	0x34, 0xae, 0xaa, 0x94, 0x93, 0xc7, 0xb3, 0xff, 0x7e, 0x33, 0xf3, 0xcf, 0xc2, 0xee, 0x25, 0x99,
	0xfc, 0xc2, 0x92, 0xb1, 0x17, 0x8d, 0x56, 0x56, 0x1d, 0xf0, 0xfe, 0x83, 0xd0, 0xa5, 0xfa, 0xd0,
	0x84, 0xfe, 0x9d, 0x66, 0xe6, 0x4e, 0xc3, 0xbd, 0xe6, 0xaa, 0x88, 0x07, 0x37, 0xaf, 0x0a, 0x77,
	0xb0, 0xff, 0x87, 0xc1, 0xf8, 0x2b, 0x19, 0x9b, 0xd1, 0xaf, 0x96, 0x8c, 0xc5, 0x08, 0xbc, 0x95,
	0x52, 0x01, 0x9b, 0x78, 0xd1, 0xab, 0x64, 0x97, 0x6f, 0xa0, 0xfc, 0xb4, 0xaf, 0xb6, 0x90, 0x6d,
	0x9d, 0x75, 0x12, 0xf4, 0xc1, 0x5b, 0x0a, 0x1d, 0x8c, 0x26, 0x2c, 0x7a, 0x91, 0x75, 0x21, 0xbe,
	0xeb, 0x32, 0x37, 0x81, 0x37, 0x61, 0xd1, 0x38, 0xd9, 0x1b, 0xde, 0xed, 0x2a, 0x7c, 0x26, 0x63,
	0x44, 0x41, 0x9d, 0xf4, 0x06, 0x4f, 0xe1, 0x99, 0x91, 0x62, 0xd5, 0x06, 0x4f, 0x7b, 0xf1, 0x87,
	0x2d, 0x62, 0x9e, 0x92, 0xb1, 0x74, 0xf9, 0xef, 0x6f, 0x2e, 0x95, 0x2d, 0x49, 0xdf, 0x4b, 0x66,
	0x0e, 0xb4, 0xff, 0x11, 0x5e, 0xba, 0x39, 0x4c, 0xa3, 0xa4, 0x21, 0x9c, 0x82, 0x27, 0x6c, 0x1d,
	0xb0, 0x9e, 0xff, 0x76, 0xc8, 0xbf, 0xa5, 0xdc, 0xeb, 0x49, 0xd8, 0xfa, 0xfd, 0x0c, 0x60, 0x33,
	0x23, 0x8e, 0xe1, 0xf9, 0x79, 0x7a, 0x92, 0x7e, 0xf9, 0x9e, 0xfa, 0x4f, 0x10, 0x60, 0xe7, 0xdb,
	0xfc, 0xd3, 0xf9, 0x62, 0xe6, 0xb3, 0xbb, 0x38, 0xf1, 0x47, 0xc9, 0xef, 0x91, 0x73, 0xef, 0x8c,
	0xf4, 0x75, 0x95, 0x13, 0x1e, 0xc1, 0xf8, 0x48, 0x9d, 0xa9, 0x9a, 0x6c, 0x59, 0xc9, 0x02, 0x1f,
	0x78, 0x70, 0xeb, 0x72, 0xf8, 0x86, 0xff, 0x2c, 0xdb, 0xba, 0x71, 0x3b, 0xd0, 0xb4, 0x5a, 0x53,
	0x6e, 0x79, 0xb7, 0x1e, 0x7e, 0x28, 0x34, 0x1e, 0xc3, 0xeb, 0x01, 0x65, 0xb1, 0x36, 0x84, 0xdb,
	0xdc, 0x0c, 0x83, 0x87, 0x25, 0x9c, 0x01, 0x11, 0xc3, 0x0c, 0xfc, 0x01, 0x67, 0x5e, 0x88, 0x4a,
	0xe2, 0xa3, 0x95, 0xc3, 0xff, 0xf8, 0x34, 0x65, 0x78, 0x02, 0x38, 0x60, 0x1e, 0x2b, 0x4d, 0xd7,
	0xa4, 0xb7, 0x0f, 0xfa, 0x48, 0x7b, 0x53, 0x76, 0x78, 0xf0, 0x63, 0x56, 0x54, 0xb6, 0x6c, 0x97,
	0x3c, 0x57, 0x75, 0xdc, 0x37, 0x16, 0x0f, 0x1b, 0x8b, 0x2b, 0x69, 0x49, 0x4b, 0xb1, 0x8e, 0x37,
	0x94, 0xe5, 0x8e, 0x7b, 0xf6, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x40, 0x1c, 0xe8, 0x60, 0x09,
	0x03, 0x00, 0x00,
}
