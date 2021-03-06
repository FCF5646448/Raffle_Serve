// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package example_v1

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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloMessage) Reset()         { *m = HelloMessage{} }
func (m *HelloMessage) String() string { return proto.CompactTextString(m) }
func (*HelloMessage) ProtoMessage()    {}
func (*HelloMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *HelloMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloMessage.Unmarshal(m, b)
}
func (m *HelloMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloMessage.Marshal(b, m, deterministic)
}
func (m *HelloMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloMessage.Merge(m, src)
}
func (m *HelloMessage) XXX_Size() int {
	return xxx_messageInfo_HelloMessage.Size(m)
}
func (m *HelloMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HelloMessage proto.InternalMessageInfo

func (m *HelloMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HelloResponse struct {
	Code                 int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *HelloMessage `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HelloResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *HelloResponse) GetData() *HelloMessage {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "example.v1.HelloRequest")
	proto.RegisterType((*HelloMessage)(nil), "example.v1.HelloMessage")
	proto.RegisterType((*HelloResponse)(nil), "example.v1.HelloResponse")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5,
	0x2b, 0x33, 0x54, 0x52, 0xe2, 0xe2, 0xf1, 0x48, 0xcd, 0xc9, 0xc9, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x95, 0x34, 0xa0, 0x6a, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x24,
	0xb8, 0xd8, 0x73, 0x21, 0x4c, 0xa8, 0x32, 0x18, 0x57, 0x29, 0x99, 0x8b, 0x17, 0x6a, 0x5a, 0x71,
	0x41, 0x7e, 0x5e, 0x71, 0x2a, 0xc8, 0xb8, 0xe4, 0xfc, 0x14, 0x88, 0x3a, 0xd6, 0x20, 0x30, 0x5b,
	0x48, 0x80, 0x8b, 0x39, 0xb7, 0x38, 0x5d, 0x82, 0x09, 0xac, 0x15, 0xc4, 0x14, 0xd2, 0xe1, 0x62,
	0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd0, 0x43, 0xb8, 0x4f,
	0x0f, 0xd9, 0xe2, 0x20, 0xb0, 0x2a, 0x23, 0x77, 0x2e, 0x16, 0xd7, 0xe4, 0x8c, 0x7c, 0x21, 0x7b,
	0x2e, 0x8e, 0xe0, 0xc4, 0x4a, 0xb0, 0x02, 0x21, 0x4c, 0x3d, 0x50, 0x0f, 0x49, 0x49, 0x62, 0x91,
	0x81, 0x38, 0x2e, 0x89, 0x0d, 0x1c, 0x1c, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xa3,
	0xd9, 0x21, 0x1c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/example.v1.Echo/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (*UnimplementedEchoServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.v1.Echo/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.v1.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Echo_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
