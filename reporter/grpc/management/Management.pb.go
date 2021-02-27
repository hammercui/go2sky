// Code generated by protoc-gen-go. DO NOT EDIT.
// source: management/Management.proto

package management

import (
	context "context"
	fmt "fmt"
	common "github.com/hammercui/go2sky/reporter/grpc/common"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type InstanceProperties struct {
	Service              string                       `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ServiceInstance      string                       `protobuf:"bytes,2,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
	Properties           []*common.KeyStringValuePair `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *InstanceProperties) Reset()         { *m = InstanceProperties{} }
func (m *InstanceProperties) String() string { return proto.CompactTextString(m) }
func (*InstanceProperties) ProtoMessage()    {}
func (*InstanceProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0d74f2a76a8a09, []int{0}
}

func (m *InstanceProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceProperties.Unmarshal(m, b)
}
func (m *InstanceProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceProperties.Marshal(b, m, deterministic)
}
func (m *InstanceProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceProperties.Merge(m, src)
}
func (m *InstanceProperties) XXX_Size() int {
	return xxx_messageInfo_InstanceProperties.Size(m)
}
func (m *InstanceProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceProperties.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceProperties proto.InternalMessageInfo

func (m *InstanceProperties) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *InstanceProperties) GetServiceInstance() string {
	if m != nil {
		return m.ServiceInstance
	}
	return ""
}

func (m *InstanceProperties) GetProperties() []*common.KeyStringValuePair {
	if m != nil {
		return m.Properties
	}
	return nil
}

type InstancePingPkg struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ServiceInstance      string   `protobuf:"bytes,2,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstancePingPkg) Reset()         { *m = InstancePingPkg{} }
func (m *InstancePingPkg) String() string { return proto.CompactTextString(m) }
func (*InstancePingPkg) ProtoMessage()    {}
func (*InstancePingPkg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0d74f2a76a8a09, []int{1}
}

func (m *InstancePingPkg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstancePingPkg.Unmarshal(m, b)
}
func (m *InstancePingPkg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstancePingPkg.Marshal(b, m, deterministic)
}
func (m *InstancePingPkg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstancePingPkg.Merge(m, src)
}
func (m *InstancePingPkg) XXX_Size() int {
	return xxx_messageInfo_InstancePingPkg.Size(m)
}
func (m *InstancePingPkg) XXX_DiscardUnknown() {
	xxx_messageInfo_InstancePingPkg.DiscardUnknown(m)
}

var xxx_messageInfo_InstancePingPkg proto.InternalMessageInfo

func (m *InstancePingPkg) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *InstancePingPkg) GetServiceInstance() string {
	if m != nil {
		return m.ServiceInstance
	}
	return ""
}

func init() {
	proto.RegisterType((*InstanceProperties)(nil), "InstanceProperties")
	proto.RegisterType((*InstancePingPkg)(nil), "InstancePingPkg")
}

func init() { proto.RegisterFile("management/Management.proto", fileDescriptor_1d0d74f2a76a8a09) }

var fileDescriptor_1d0d74f2a76a8a09 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x51, 0x3d, 0x4b, 0x03, 0x41,
	0x14, 0xf4, 0x12, 0x50, 0xb2, 0x16, 0xd1, 0x4b, 0x73, 0xc4, 0x26, 0xa4, 0x0a, 0x16, 0xbb, 0x98,
	0xb4, 0x36, 0xd1, 0x4a, 0x24, 0x72, 0xe4, 0x50, 0xc1, 0x6e, 0xb3, 0x3e, 0x36, 0xcb, 0xde, 0x7e,
	0xf0, 0x76, 0x93, 0x70, 0x85, 0x95, 0xb5, 0x7f, 0xc4, 0x5f, 0x29, 0xf9, 0xd6, 0x68, 0x69, 0xb5,
	0x6f, 0x87, 0x99, 0x61, 0x98, 0x21, 0x17, 0x86, 0x5b, 0x2e, 0xc1, 0x80, 0x8d, 0x6c, 0xb4, 0x3b,
	0xa9, 0x47, 0x17, 0x5d, 0xbb, 0x25, 0x9c, 0x31, 0xce, 0xb2, 0xdb, 0xd5, 0xb3, 0x06, 0xbb, 0x1f,
	0x09, 0x49, 0xef, 0x6c, 0x88, 0xdc, 0x0a, 0xc8, 0xd1, 0x79, 0xc0, 0xa8, 0x20, 0xa4, 0x19, 0x39,
	0x09, 0x80, 0x73, 0x25, 0x20, 0x4b, 0x3a, 0x49, 0xaf, 0x31, 0xde, 0x7e, 0xd3, 0x1e, 0x69, 0x6e,
	0xce, 0xad, 0x2c, 0xab, 0xad, 0x18, 0x87, 0x70, 0x3a, 0x20, 0xc4, 0xef, 0x1c, 0xb3, 0x7a, 0xa7,
	0xde, 0x3b, 0xed, 0xb7, 0xe8, 0x3d, 0x54, 0x45, 0x44, 0x65, 0xe5, 0x13, 0x2f, 0x67, 0x90, 0x73,
	0x85, 0xe3, 0x6f, 0xb4, 0xee, 0x23, 0x69, 0xee, 0xe2, 0x28, 0x2b, 0x73, 0x2d, 0xff, 0x23, 0x4b,
	0xff, 0x8d, 0x9c, 0xef, 0xfb, 0x28, 0x36, 0xf2, 0x6b, 0x92, 0x21, 0x78, 0x87, 0xf1, 0x8f, 0x02,
	0x5a, 0xf4, 0x37, 0xd8, 0x6e, 0xd0, 0x65, 0x77, 0xdc, 0xbe, 0x86, 0xee, 0x51, 0x7a, 0x49, 0x1a,
	0x1a, 0xc0, 0x0f, 0x4b, 0x35, 0x87, 0xf4, 0x8c, 0x1e, 0xa4, 0xfe, 0xc1, 0xbd, 0x79, 0x4f, 0x08,
	0x73, 0x28, 0x29, 0xf7, 0x5c, 0x4c, 0x81, 0x06, 0x5d, 0x2d, 0x78, 0xa9, 0x95, 0x5d, 0x22, 0x86,
	0x5a, 0x88, 0x0b, 0x87, 0x9a, 0xee, 0xc7, 0xa3, 0xf3, 0x41, 0x9e, 0xbc, 0x5c, 0x49, 0x15, 0xa7,
	0xb3, 0x09, 0x15, 0xce, 0xb0, 0x42, 0x57, 0xc3, 0x7c, 0xc4, 0xa4, 0xeb, 0x07, 0x5d, 0xb1, 0x75,
	0x6c, 0x40, 0x26, 0xd1, 0x0b, 0xb6, 0xd7, 0x7d, 0xd6, 0xda, 0x85, 0xae, 0x9e, 0x37, 0xde, 0x0f,
	0x6b, 0xdf, 0x7c, 0x39, 0xb4, 0x70, 0xe5, 0xe4, 0x78, 0x35, 0xf9, 0xe0, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0xbb, 0x3d, 0xc3, 0x47, 0x26, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ManagementServiceClient is the client API for ManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManagementServiceClient interface {
	ReportInstanceProperties(ctx context.Context, in *InstanceProperties, opts ...grpc.CallOption) (*common.Commands, error)
	KeepAlive(ctx context.Context, in *InstancePingPkg, opts ...grpc.CallOption) (*common.Commands, error)
}

type managementServiceClient struct {
	cc *grpc.ClientConn
}

func NewManagementServiceClient(cc *grpc.ClientConn) ManagementServiceClient {
	return &managementServiceClient{cc}
}

func (c *managementServiceClient) ReportInstanceProperties(ctx context.Context, in *InstanceProperties, opts ...grpc.CallOption) (*common.Commands, error) {
	out := new(common.Commands)
	err := c.cc.Invoke(ctx, "/ManagementService/reportInstanceProperties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementServiceClient) KeepAlive(ctx context.Context, in *InstancePingPkg, opts ...grpc.CallOption) (*common.Commands, error) {
	out := new(common.Commands)
	err := c.cc.Invoke(ctx, "/ManagementService/keepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagementServiceServer is the server API for ManagementService service.
type ManagementServiceServer interface {
	ReportInstanceProperties(context.Context, *InstanceProperties) (*common.Commands, error)
	KeepAlive(context.Context, *InstancePingPkg) (*common.Commands, error)
}

func RegisterManagementServiceServer(s *grpc.Server, srv ManagementServiceServer) {
	s.RegisterService(&_ManagementService_serviceDesc, srv)
}

func _ManagementService_ReportInstanceProperties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstanceProperties)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).ReportInstanceProperties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ManagementService/ReportInstanceProperties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).ReportInstanceProperties(ctx, req.(*InstanceProperties))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagementService_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstancePingPkg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServiceServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ManagementService/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServiceServer).KeepAlive(ctx, req.(*InstancePingPkg))
	}
	return interceptor(ctx, in, info, handler)
}

var _ManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ManagementService",
	HandlerType: (*ManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "reportInstanceProperties",
			Handler:    _ManagementService_ReportInstanceProperties_Handler,
		},
		{
			MethodName: "keepAlive",
			Handler:    _ManagementService_KeepAlive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "management/Management.proto",
}
