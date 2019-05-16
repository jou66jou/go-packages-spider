// Code generated by protoc-gen-go. DO NOT EDIT.
// source: route_pkgs.proto

package echo

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//定義 EchoPkg Request
type GOPkgRequest struct {
	PkgName              string   `protobuf:"bytes,1,opt,name=pkg_name,json=pkgName,proto3" json:"pkg_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GOPkgRequest) Reset()         { *m = GOPkgRequest{} }
func (m *GOPkgRequest) String() string { return proto.CompactTextString(m) }
func (*GOPkgRequest) ProtoMessage()    {}
func (*GOPkgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d07b2d409cbda4f1, []int{0}
}

func (m *GOPkgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GOPkgRequest.Unmarshal(m, b)
}
func (m *GOPkgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GOPkgRequest.Marshal(b, m, deterministic)
}
func (m *GOPkgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GOPkgRequest.Merge(m, src)
}
func (m *GOPkgRequest) XXX_Size() int {
	return xxx_messageInfo_GOPkgRequest.Size(m)
}
func (m *GOPkgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GOPkgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GOPkgRequest proto.InternalMessageInfo

func (m *GOPkgRequest) GetPkgName() string {
	if m != nil {
		return m.PkgName
	}
	return ""
}

//定義 EchoPkg Response
type GOPkgReply struct {
	PkgName              string   `protobuf:"bytes,1,opt,name=pkg_name,json=pkgName,proto3" json:"pkg_name,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Overview             string   `protobuf:"bytes,3,opt,name=overview,proto3" json:"overview,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GOPkgReply) Reset()         { *m = GOPkgReply{} }
func (m *GOPkgReply) String() string { return proto.CompactTextString(m) }
func (*GOPkgReply) ProtoMessage()    {}
func (*GOPkgReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d07b2d409cbda4f1, []int{1}
}

func (m *GOPkgReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GOPkgReply.Unmarshal(m, b)
}
func (m *GOPkgReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GOPkgReply.Marshal(b, m, deterministic)
}
func (m *GOPkgReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GOPkgReply.Merge(m, src)
}
func (m *GOPkgReply) XXX_Size() int {
	return xxx_messageInfo_GOPkgReply.Size(m)
}
func (m *GOPkgReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GOPkgReply.DiscardUnknown(m)
}

var xxx_messageInfo_GOPkgReply proto.InternalMessageInfo

func (m *GOPkgReply) GetPkgName() string {
	if m != nil {
		return m.PkgName
	}
	return ""
}

func (m *GOPkgReply) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *GOPkgReply) GetOverview() string {
	if m != nil {
		return m.Overview
	}
	return ""
}

func init() {
	proto.RegisterType((*GOPkgRequest)(nil), "echo.GOPkgRequest")
	proto.RegisterType((*GOPkgReply)(nil), "echo.GOPkgReply")
}

func init() { proto.RegisterFile("route_pkgs.proto", fileDescriptor_d07b2d409cbda4f1) }

var fileDescriptor_d07b2d409cbda4f1 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xca, 0x2f, 0x2d,
	0x49, 0x8d, 0x2f, 0xc8, 0x4e, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x4d,
	0xce, 0xc8, 0x57, 0xd2, 0xe4, 0xe2, 0x71, 0xf7, 0x0f, 0xc8, 0x4e, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x92, 0xe4, 0xe2, 0x28, 0xc8, 0x4e, 0x8f, 0xcf, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0x2f, 0xc8, 0x4e, 0xf7, 0x4b, 0xcc, 0x4d, 0x55, 0x0a, 0xe5,
	0xe2, 0x82, 0x2a, 0x2d, 0xc8, 0xa9, 0xc4, 0xa3, 0x50, 0x48, 0x80, 0x8b, 0xb9, 0xb4, 0x28, 0x47,
	0x82, 0x09, 0x2c, 0x0a, 0x62, 0x0a, 0x49, 0x71, 0x71, 0xe4, 0x97, 0xa5, 0x16, 0x95, 0x65, 0xa6,
	0x96, 0x4b, 0x30, 0x83, 0x85, 0xe1, 0x7c, 0x23, 0x1b, 0x2e, 0x76, 0xd7, 0xe4, 0x8c, 0xfc, 0x80,
	0xec, 0x74, 0x21, 0x43, 0x04, 0x53, 0x48, 0x0f, 0xe4, 0x3c, 0x3d, 0x64, 0xb7, 0x49, 0x09, 0xa0,
	0x88, 0x15, 0xe4, 0x54, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0x3d, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x3d, 0x7c, 0x20, 0xc8, 0xe0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoPkgClient is the client API for EchoPkg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoPkgClient interface {
	EchoPkg(ctx context.Context, in *GOPkgRequest, opts ...grpc.CallOption) (*GOPkgReply, error)
}

type echoPkgClient struct {
	cc *grpc.ClientConn
}

func NewEchoPkgClient(cc *grpc.ClientConn) EchoPkgClient {
	return &echoPkgClient{cc}
}

func (c *echoPkgClient) EchoPkg(ctx context.Context, in *GOPkgRequest, opts ...grpc.CallOption) (*GOPkgReply, error) {
	out := new(GOPkgReply)
	err := c.cc.Invoke(ctx, "/echo.EchoPkg/EchoPkg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoPkgServer is the server API for EchoPkg service.
type EchoPkgServer interface {
	EchoPkg(context.Context, *GOPkgRequest) (*GOPkgReply, error)
}

// UnimplementedEchoPkgServer can be embedded to have forward compatible implementations.
type UnimplementedEchoPkgServer struct {
}

func (*UnimplementedEchoPkgServer) EchoPkg(ctx context.Context, req *GOPkgRequest) (*GOPkgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoPkg not implemented")
}

func RegisterEchoPkgServer(s *grpc.Server, srv EchoPkgServer) {
	s.RegisterService(&_EchoPkg_serviceDesc, srv)
}

func _EchoPkg_EchoPkg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GOPkgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoPkgServer).EchoPkg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.EchoPkg/EchoPkg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoPkgServer).EchoPkg(ctx, req.(*GOPkgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoPkg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.EchoPkg",
	HandlerType: (*EchoPkgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EchoPkg",
			Handler:    _EchoPkg_EchoPkg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "route_pkgs.proto",
}