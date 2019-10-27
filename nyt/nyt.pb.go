// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nyt.proto

package nyt

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

type NytRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Period               int32    `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NytRequest) Reset()         { *m = NytRequest{} }
func (m *NytRequest) String() string { return proto.CompactTextString(m) }
func (*NytRequest) ProtoMessage()    {}
func (*NytRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27bdeb0e5bc92888, []int{0}
}

func (m *NytRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NytRequest.Unmarshal(m, b)
}
func (m *NytRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NytRequest.Marshal(b, m, deterministic)
}
func (m *NytRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NytRequest.Merge(m, src)
}
func (m *NytRequest) XXX_Size() int {
	return xxx_messageInfo_NytRequest.Size(m)
}
func (m *NytRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NytRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NytRequest proto.InternalMessageInfo

func (m *NytRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *NytRequest) GetPeriod() int32 {
	if m != nil {
		return m.Period
	}
	return 0
}

type NytResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NytResponse) Reset()         { *m = NytResponse{} }
func (m *NytResponse) String() string { return proto.CompactTextString(m) }
func (*NytResponse) ProtoMessage()    {}
func (*NytResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27bdeb0e5bc92888, []int{1}
}

func (m *NytResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NytResponse.Unmarshal(m, b)
}
func (m *NytResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NytResponse.Marshal(b, m, deterministic)
}
func (m *NytResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NytResponse.Merge(m, src)
}
func (m *NytResponse) XXX_Size() int {
	return xxx_messageInfo_NytResponse.Size(m)
}
func (m *NytResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NytResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NytResponse proto.InternalMessageInfo

func (m *NytResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*NytRequest)(nil), "nyt.NytRequest")
	proto.RegisterType((*NytResponse)(nil), "nyt.NytResponse")
}

func init() { proto.RegisterFile("nyt.proto", fileDescriptor_27bdeb0e5bc92888) }

var fileDescriptor_27bdeb0e5bc92888 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0xab, 0x2c, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0xab, 0x2c, 0x51, 0xb2, 0xe2, 0xe2, 0xf2, 0xab,
	0x2c, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0xcf, 0x4e,
	0xcd, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0xc4, 0xb8, 0xd8, 0x0a, 0x52,
	0x8b, 0x32, 0xf3, 0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c, 0x25, 0x4d, 0x2e,
	0x6e, 0xb0, 0xde, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x29, 0x2e, 0x8e, 0x22, 0x28, 0x1b,
	0xaa, 0x1f, 0xce, 0x37, 0x72, 0x02, 0x5b, 0x13, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64,
	0xc2, 0xc5, 0xeb, 0x9e, 0x5a, 0xe2, 0x9b, 0x5f, 0x5c, 0x12, 0x9c, 0x91, 0x58, 0x94, 0x9a, 0x22,
	0xc4, 0xaf, 0x07, 0x72, 0x16, 0xc2, 0x21, 0x52, 0x02, 0x08, 0x01, 0x88, 0x09, 0x4a, 0x0c, 0x49,
	0x6c, 0x60, 0x67, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x47, 0x40, 0x8d, 0xc3, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NytServiceClient is the client API for NytService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NytServiceClient interface {
	GetMostShared(ctx context.Context, in *NytRequest, opts ...grpc.CallOption) (*NytResponse, error)
}

type nytServiceClient struct {
	cc *grpc.ClientConn
}

func NewNytServiceClient(cc *grpc.ClientConn) NytServiceClient {
	return &nytServiceClient{cc}
}

func (c *nytServiceClient) GetMostShared(ctx context.Context, in *NytRequest, opts ...grpc.CallOption) (*NytResponse, error) {
	out := new(NytResponse)
	err := c.cc.Invoke(ctx, "/nyt.NytService/GetMostShared", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NytServiceServer is the server API for NytService service.
type NytServiceServer interface {
	GetMostShared(context.Context, *NytRequest) (*NytResponse, error)
}

// UnimplementedNytServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNytServiceServer struct {
}

func (*UnimplementedNytServiceServer) GetMostShared(ctx context.Context, req *NytRequest) (*NytResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMostShared not implemented")
}

func RegisterNytServiceServer(s *grpc.Server, srv NytServiceServer) {
	s.RegisterService(&_NytService_serviceDesc, srv)
}

func _NytService_GetMostShared_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NytRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NytServiceServer).GetMostShared(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nyt.NytService/GetMostShared",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NytServiceServer).GetMostShared(ctx, req.(*NytRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NytService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nyt.NytService",
	HandlerType: (*NytServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMostShared",
			Handler:    _NytService_GetMostShared_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nyt.proto",
}