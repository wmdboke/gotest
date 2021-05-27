// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Req struct {
	Fromjson             string   `protobuf:"bytes,1,opt,name=fromjson,proto3" json:"fromjson,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetFromjson() string {
	if m != nil {
		return m.Fromjson
	}
	return ""
}

type Rep struct {
	Md5Str               string   `protobuf:"bytes,1,opt,name=md5str,proto3" json:"md5str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rep) Reset()         { *m = Rep{} }
func (m *Rep) String() string { return proto.CompactTextString(m) }
func (*Rep) ProtoMessage()    {}
func (*Rep) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *Rep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rep.Unmarshal(m, b)
}
func (m *Rep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rep.Marshal(b, m, deterministic)
}
func (m *Rep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rep.Merge(m, src)
}
func (m *Rep) XXX_Size() int {
	return xxx_messageInfo_Rep.Size(m)
}
func (m *Rep) XXX_DiscardUnknown() {
	xxx_messageInfo_Rep.DiscardUnknown(m)
}

var xxx_messageInfo_Rep proto.InternalMessageInfo

func (m *Rep) GetMd5Str() string {
	if m != nil {
		return m.Md5Str
	}
	return ""
}

func init() {
	proto.RegisterType((*Req)(nil), "test.Req")
	proto.RegisterType((*Rep)(nil), "test.Rep")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x14, 0xb9, 0x98, 0x83, 0x52,
	0x0b, 0x85, 0xa4, 0xb8, 0x38, 0xd2, 0x8a, 0xf2, 0x73, 0xb3, 0x8a, 0xf3, 0xf3, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x25, 0x59, 0x90, 0x92, 0x02, 0x21, 0x31, 0x2e, 0xb6, 0xdc,
	0x14, 0xd3, 0xe2, 0x92, 0x22, 0xa8, 0x02, 0x28, 0xcf, 0x48, 0x9d, 0x8b, 0x25, 0x3c, 0x31, 0xb3,
	0x44, 0x48, 0x9e, 0x8b, 0xd5, 0x25, 0xdf, 0xd7, 0xc5, 0x54, 0x88, 0x53, 0x0f, 0x6c, 0x4b, 0x50,
	0x6a, 0xa1, 0x14, 0x9c, 0x59, 0xa0, 0xc4, 0x90, 0xc4, 0x06, 0xb6, 0xd7, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x8d, 0x68, 0x98, 0xa0, 0x85, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WaitClient is the client API for Wait service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WaitClient interface {
	DoMD5(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rep, error)
}

type waitClient struct {
	cc *grpc.ClientConn
}

func NewWaitClient(cc *grpc.ClientConn) WaitClient {
	return &waitClient{cc}
}

func (c *waitClient) DoMD5(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rep, error) {
	out := new(Rep)
	err := c.cc.Invoke(ctx, "/test.Wait/DoMD5", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WaitServer is the server API for Wait service.
type WaitServer interface {
	DoMD5(context.Context, *Req) (*Rep, error)
}

// UnimplementedWaitServer can be embedded to have forward compatible implementations.
type UnimplementedWaitServer struct {
}

func (*UnimplementedWaitServer) DoMD5(ctx context.Context, req *Req) (*Rep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoMD5 not implemented")
}

func RegisterWaitServer(s *grpc.Server, srv WaitServer) {
	s.RegisterService(&_Wait_serviceDesc, srv)
}

func _Wait_DoMD5_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaitServer).DoMD5(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Wait/DoMD5",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaitServer).DoMD5(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _Wait_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.Wait",
	HandlerType: (*WaitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoMD5",
			Handler:    _Wait_DoMD5_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
