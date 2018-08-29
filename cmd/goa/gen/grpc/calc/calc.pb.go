// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gen/grpc/calc/calc.proto

/*
Package calc is a generated protocol buffer package.

It is generated from these files:
	gen/grpc/calc/calc.proto

It has these top-level messages:
	AddRequest
	AddResponse
*/
package calc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddRequest struct {
	A int32               `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int32               `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
	C *AddRequest_FooType `protobuf:"bytes,3,opt,name=c" json:"c,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *AddRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *AddRequest) GetC() *AddRequest_FooType {
	if m != nil {
		return m.C
	}
	return nil
}

// Sum
type AddRequest_FooType struct {
	C int32 `protobuf:"varint,1,opt,name=c" json:"c,omitempty"`
}

func (m *AddRequest_FooType) Reset()                    { *m = AddRequest_FooType{} }
func (m *AddRequest_FooType) String() string            { return proto.CompactTextString(m) }
func (*AddRequest_FooType) ProtoMessage()               {}
func (*AddRequest_FooType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *AddRequest_FooType) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

type AddResponse struct {
	Field int32 `protobuf:"varint,1,opt,name=field" json:"field,omitempty"`
}

func (m *AddResponse) Reset()                    { *m = AddResponse{} }
func (m *AddResponse) String() string            { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()               {}
func (*AddResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddResponse) GetField() int32 {
	if m != nil {
		return m.Field
	}
	return 0
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "calc.AddRequest")
	proto.RegisterType((*AddRequest_FooType)(nil), "calc.AddRequest.FooType")
	proto.RegisterType((*AddResponse)(nil), "calc.AddResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Calc service

type CalcClient interface {
	// Adds the operands
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type calcClient struct {
	cc *grpc.ClientConn
}

func NewCalcClient(cc *grpc.ClientConn) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := grpc.Invoke(ctx, "/calc.calc/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Calc service

type CalcServer interface {
	// Adds the operands
	Add(context.Context, *AddRequest) (*AddResponse, error)
}

func RegisterCalcServer(s *grpc.Server, srv CalcServer) {
	s.RegisterService(&_Calc_serviceDesc, srv)
}

func _Calc_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.calc/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.calc",
	HandlerType: (*CalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calc_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gen/grpc/calc/calc.proto",
}

func init() { proto.RegisterFile("gen/grpc/calc/calc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x4f, 0xcd, 0xd3,
	0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x4f, 0x4e, 0xcc, 0x81, 0x10, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x2c, 0x20, 0xb6, 0x52, 0x26, 0x17, 0x97, 0x63, 0x4a, 0x4a, 0x50, 0x6a, 0x61, 0x69, 0x6a,
	0x71, 0x89, 0x10, 0x0f, 0x17, 0x63, 0xa2, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x63, 0x22,
	0x88, 0x97, 0x24, 0xc1, 0x04, 0xe1, 0x25, 0x09, 0xa9, 0x71, 0x31, 0x26, 0x4b, 0x30, 0x2b, 0x30,
	0x6a, 0x70, 0x1b, 0x49, 0xe8, 0x81, 0xcd, 0x41, 0x68, 0xd4, 0x73, 0xcb, 0xcf, 0x0f, 0xa9, 0x2c,
	0x48, 0x0d, 0x62, 0x4c, 0x96, 0x12, 0xe7, 0x62, 0x87, 0xf2, 0x40, 0x06, 0x24, 0xc3, 0x8c, 0x4b,
	0x56, 0x52, 0xe6, 0xe2, 0x06, 0xeb, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe1, 0x62,
	0x4d, 0xcb, 0x4c, 0xcd, 0x49, 0x81, 0x2a, 0x80, 0x70, 0x8c, 0x4c, 0xb8, 0xc0, 0xee, 0x12, 0xd2,
	0xe1, 0x62, 0x76, 0x4c, 0x49, 0x11, 0x12, 0x40, 0xb7, 0x49, 0x4a, 0x10, 0x49, 0x04, 0x62, 0x92,
	0x12, 0x43, 0x12, 0x1b, 0xd8, 0x4b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x3a, 0x52,
	0xec, 0xee, 0x00, 0x00, 0x00,
}