// Code generated by protoc-gen-go. DO NOT EDIT.
// source: squareroot/squarerootpb/squareroot.proto

package squarerootpb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type SquareRootRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootRequest) Reset()         { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()    {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e9636629a13a5e0, []int{0}
}

func (m *SquareRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootRequest.Unmarshal(m, b)
}
func (m *SquareRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootRequest.Marshal(b, m, deterministic)
}
func (m *SquareRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootRequest.Merge(m, src)
}
func (m *SquareRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRootRequest.Size(m)
}
func (m *SquareRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootRequest proto.InternalMessageInfo

func (m *SquareRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareRootResponse struct {
	Root                 float64  `protobuf:"fixed64,1,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootResponse) Reset()         { *m = SquareRootResponse{} }
func (m *SquareRootResponse) String() string { return proto.CompactTextString(m) }
func (*SquareRootResponse) ProtoMessage()    {}
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e9636629a13a5e0, []int{1}
}

func (m *SquareRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootResponse.Unmarshal(m, b)
}
func (m *SquareRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootResponse.Marshal(b, m, deterministic)
}
func (m *SquareRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootResponse.Merge(m, src)
}
func (m *SquareRootResponse) XXX_Size() int {
	return xxx_messageInfo_SquareRootResponse.Size(m)
}
func (m *SquareRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootResponse proto.InternalMessageInfo

func (m *SquareRootResponse) GetRoot() float64 {
	if m != nil {
		return m.Root
	}
	return 0
}

func init() {
	proto.RegisterType((*SquareRootRequest)(nil), "squareroot.SquareRootRequest")
	proto.RegisterType((*SquareRootResponse)(nil), "squareroot.SquareRootResponse")
}

func init() {
	proto.RegisterFile("squareroot/squarerootpb/squareroot.proto", fileDescriptor_4e9636629a13a5e0)
}

var fileDescriptor_4e9636629a13a5e0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x2e, 0x2c, 0x4d,
	0x2c, 0x4a, 0x2d, 0xca, 0xcf, 0x2f, 0xd1, 0x47, 0x30, 0x0b, 0x92, 0x90, 0x38, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x5c, 0x08, 0x11, 0x25, 0x6d, 0x2e, 0xc1, 0x60, 0x30, 0x2f, 0x28, 0x3f,
	0xbf, 0x24, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x8c, 0x8b, 0x2d, 0xaf, 0x34, 0x37,
	0x29, 0xb5, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x08, 0xca, 0x53, 0xd2, 0xe0, 0x12, 0x42,
	0x56, 0x5c, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc4, 0xc5, 0x02, 0x32, 0x0a, 0xac, 0x96,
	0x31, 0x08, 0xcc, 0x36, 0x4a, 0x42, 0x36, 0x36, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8,
	0x97, 0x8b, 0x0b, 0x21, 0x28, 0x24, 0xab, 0x87, 0xe4, 0x30, 0x0c, 0x37, 0x48, 0xc9, 0xe1, 0x92,
	0x86, 0xd8, 0xaa, 0xc4, 0xe0, 0xc4, 0x17, 0xc5, 0x83, 0xec, 0xcf, 0x24, 0x36, 0xb0, 0xef, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x43, 0x9c, 0x1b, 0x09, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SquareRootServiceClient is the client API for SquareRootService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SquareRootServiceClient interface {
	// Error handling
	// This RPC will throw an exception if the sent number is negative
	// The error being sent is of type INVALID_ARGUMENT
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type squareRootServiceClient struct {
	cc *grpc.ClientConn
}

func NewSquareRootServiceClient(cc *grpc.ClientConn) SquareRootServiceClient {
	return &squareRootServiceClient{cc}
}

func (c *squareRootServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/squareroot.SquareRootService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SquareRootServiceServer is the server API for SquareRootService service.
type SquareRootServiceServer interface {
	// Error handling
	// This RPC will throw an exception if the sent number is negative
	// The error being sent is of type INVALID_ARGUMENT
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

func RegisterSquareRootServiceServer(s *grpc.Server, srv SquareRootServiceServer) {
	s.RegisterService(&_SquareRootService_serviceDesc, srv)
}

func _SquareRootService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquareRootServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squareroot.SquareRootService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquareRootServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SquareRootService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "squareroot.SquareRootService",
	HandlerType: (*SquareRootServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SquareRoot",
			Handler:    _SquareRootService_SquareRoot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "squareroot/squarerootpb/squareroot.proto",
}
