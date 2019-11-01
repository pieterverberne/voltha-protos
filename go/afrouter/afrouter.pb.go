// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/afrouter.proto

package afrouter

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/opencord/voltha-protos/v2/go/common"
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

type Result struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Info                 string   `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Result) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Count struct {
	Count                uint32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Count) Reset()         { *m = Count{} }
func (m *Count) String() string { return proto.CompactTextString(m) }
func (*Count) ProtoMessage()    {}
func (*Count) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{2}
}

func (m *Count) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Count.Unmarshal(m, b)
}
func (m *Count) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Count.Marshal(b, m, deterministic)
}
func (m *Count) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Count.Merge(m, src)
}
func (m *Count) XXX_Size() int {
	return xxx_messageInfo_Count.Size(m)
}
func (m *Count) XXX_DiscardUnknown() {
	xxx_messageInfo_Count.DiscardUnknown(m)
}

var xxx_messageInfo_Count proto.InternalMessageInfo

func (m *Count) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Conn struct {
	Server               string   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Pkg                  string   `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	Svc                  string   `protobuf:"bytes,3,opt,name=svc,proto3" json:"svc,omitempty"`
	Cluster              string   `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Backend              string   `protobuf:"bytes,5,opt,name=backend,proto3" json:"backend,omitempty"`
	Connection           string   `protobuf:"bytes,6,opt,name=connection,proto3" json:"connection,omitempty"`
	Addr                 string   `protobuf:"bytes,7,opt,name=addr,proto3" json:"addr,omitempty"`
	Port                 uint64   `protobuf:"varint,8,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Conn) Reset()         { *m = Conn{} }
func (m *Conn) String() string { return proto.CompactTextString(m) }
func (*Conn) ProtoMessage()    {}
func (*Conn) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{3}
}

func (m *Conn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conn.Unmarshal(m, b)
}
func (m *Conn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conn.Marshal(b, m, deterministic)
}
func (m *Conn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conn.Merge(m, src)
}
func (m *Conn) XXX_Size() int {
	return xxx_messageInfo_Conn.Size(m)
}
func (m *Conn) XXX_DiscardUnknown() {
	xxx_messageInfo_Conn.DiscardUnknown(m)
}

var xxx_messageInfo_Conn proto.InternalMessageInfo

func (m *Conn) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *Conn) GetPkg() string {
	if m != nil {
		return m.Pkg
	}
	return ""
}

func (m *Conn) GetSvc() string {
	if m != nil {
		return m.Svc
	}
	return ""
}

func (m *Conn) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Conn) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func (m *Conn) GetConnection() string {
	if m != nil {
		return m.Connection
	}
	return ""
}

func (m *Conn) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Conn) GetPort() uint64 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Affinity struct {
	Router               string   `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
	Route                string   `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	Cluster              string   `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Backend              string   `protobuf:"bytes,4,opt,name=backend,proto3" json:"backend,omitempty"`
	Id                   string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Affinity) Reset()         { *m = Affinity{} }
func (m *Affinity) String() string { return proto.CompactTextString(m) }
func (*Affinity) ProtoMessage()    {}
func (*Affinity) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e2f565b9e1c96, []int{4}
}

func (m *Affinity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Affinity.Unmarshal(m, b)
}
func (m *Affinity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Affinity.Marshal(b, m, deterministic)
}
func (m *Affinity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Affinity.Merge(m, src)
}
func (m *Affinity) XXX_Size() int {
	return xxx_messageInfo_Affinity.Size(m)
}
func (m *Affinity) XXX_DiscardUnknown() {
	xxx_messageInfo_Affinity.DiscardUnknown(m)
}

var xxx_messageInfo_Affinity proto.InternalMessageInfo

func (m *Affinity) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *Affinity) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *Affinity) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Affinity) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func (m *Affinity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Result)(nil), "afrouter.Result")
	proto.RegisterType((*Empty)(nil), "afrouter.Empty")
	proto.RegisterType((*Count)(nil), "afrouter.Count")
	proto.RegisterType((*Conn)(nil), "afrouter.Conn")
	proto.RegisterType((*Affinity)(nil), "afrouter.Affinity")
}

func init() { proto.RegisterFile("voltha_protos/afrouter.proto", fileDescriptor_be7e2f565b9e1c96) }

var fileDescriptor_be7e2f565b9e1c96 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0xed, 0x23, 0x7d, 0xcc, 0x85, 0x3e, 0xb0, 0x10, 0xb2, 0x22, 0x40, 0x55, 0x56, 0xdd, 0xd0,
	0xa2, 0x0e, 0x88, 0x0d, 0x1b, 0x88, 0x50, 0x37, 0x5d, 0x65, 0xc4, 0x86, 0x0d, 0x4a, 0x1d, 0x27,
	0x63, 0x4d, 0xeb, 0x1b, 0xd9, 0x4e, 0xa4, 0x91, 0xf8, 0x20, 0xbe, 0x82, 0x6f, 0x43, 0xb6, 0x93,
	0xb6, 0x53, 0xc1, 0xee, 0x9c, 0x73, 0xef, 0xb5, 0xcf, 0xf1, 0x03, 0x5e, 0xd7, 0x78, 0x30, 0xf7,
	0xe9, 0xcf, 0x52, 0xa1, 0x41, 0xbd, 0x4e, 0x73, 0x85, 0x95, 0xe1, 0x6a, 0xe5, 0x38, 0x19, 0xb7,
	0x3c, 0x0c, 0x9f, 0xf6, 0x31, 0x3c, 0x1e, 0x51, 0xfa, 0xae, 0x68, 0x07, 0xc3, 0x84, 0xeb, 0xea,
	0x60, 0x08, 0x85, 0x91, 0xae, 0x18, 0xe3, 0x5a, 0xd3, 0xee, 0xa2, 0xbb, 0x1c, 0x27, 0x2d, 0x25,
	0x2f, 0x61, 0xc0, 0x95, 0x42, 0x45, 0x7b, 0x8b, 0xee, 0xf2, 0x26, 0xf1, 0x84, 0x10, 0x08, 0x84,
	0xcc, 0x91, 0xf6, 0x9d, 0xe8, 0x70, 0x34, 0x82, 0xc1, 0xb7, 0x63, 0x69, 0x1e, 0xa3, 0x37, 0x30,
	0x88, 0xb1, 0x92, 0xc6, 0xce, 0x32, 0x0b, 0xdc, 0x9a, 0x93, 0xc4, 0x93, 0xe8, 0x4f, 0x17, 0x82,
	0x18, 0xa5, 0x24, 0xaf, 0x60, 0xa8, 0xb9, 0xaa, 0xb9, 0x72, 0xf5, 0x9b, 0xa4, 0x61, 0x64, 0x0e,
	0xfd, 0xf2, 0xa1, 0x68, 0x36, 0xb4, 0xd0, 0x2a, 0xba, 0x66, 0xcd, 0x6e, 0x16, 0x5a, 0xc3, 0xec,
	0x50, 0x69, 0xc3, 0x15, 0x0d, 0x9c, 0xda, 0x52, 0x5b, 0xd9, 0xa7, 0xec, 0x81, 0xcb, 0x8c, 0x0e,
	0x7c, 0xa5, 0xa1, 0xe4, 0x2d, 0x00, 0x43, 0x29, 0x39, 0x33, 0x02, 0x25, 0x1d, 0xba, 0xe2, 0x85,
	0x62, 0x43, 0xa5, 0x59, 0xa6, 0xe8, 0xc8, 0x87, 0xb2, 0xd8, 0x6a, 0x25, 0x2a, 0x43, 0xc7, 0x8b,
	0xee, 0x32, 0x48, 0x1c, 0x8e, 0x7e, 0xc1, 0xf8, 0x4b, 0x9e, 0x0b, 0x29, 0xcc, 0xa3, 0xcd, 0xe0,
	0x0f, 0xba, 0xcd, 0xe0, 0x99, 0x8d, 0xee, 0x50, 0x7b, 0x6c, 0x8e, 0x5c, 0xba, 0xee, 0xff, 0xd7,
	0x75, 0xf0, 0xd4, 0xf5, 0x14, 0x7a, 0xa2, 0x8d, 0xd2, 0x13, 0xd9, 0xe6, 0x77, 0x0f, 0x26, 0x31,
	0xca, 0x5c, 0x14, 0x95, 0x4a, 0x9d, 0xef, 0x5b, 0x98, 0xdc, 0x71, 0x13, 0x9f, 0x83, 0x4c, 0x57,
	0xa7, 0xe7, 0x60, 0xd5, 0x70, 0x7e, 0xe6, 0xfe, 0xbe, 0xa3, 0x0e, 0xf9, 0x08, 0xcf, 0xee, 0xb8,
	0x39, 0xe5, 0x20, 0xe7, 0x96, 0x56, 0xfb, 0xe7, 0xd8, 0x27, 0x78, 0xb1, 0xe5, 0x66, 0x8b, 0x56,
	0x17, 0x92, 0xfb, 0x7b, 0x9e, 0x9d, 0x1b, 0xdd, 0x0b, 0x08, 0x67, 0x97, 0x06, 0xec, 0x9d, 0x77,
	0xc8, 0x07, 0x98, 0x7e, 0x2f, 0xb3, 0xd4, 0xf0, 0x1d, 0x16, 0x3b, 0x5e, 0xf3, 0x03, 0x99, 0xad,
	0x9a, 0xc7, 0xb8, 0xc3, 0xa2, 0x10, 0xb2, 0x08, 0xaf, 0x97, 0x89, 0x3a, 0xe4, 0x33, 0x3c, 0xdf,
	0x72, 0xd3, 0x8e, 0x68, 0x42, 0xaf, 0x66, 0x62, 0x3c, 0x96, 0x28, 0xb9, 0x34, 0xe1, 0xfc, 0xaa,
	0xa2, 0xa3, 0xce, 0xd7, 0xcd, 0x8f, 0xf7, 0x85, 0x30, 0xf7, 0xd5, 0xde, 0xd6, 0xd6, 0x58, 0x72,
	0xc9, 0x50, 0x65, 0x6b, 0xff, 0x23, 0xde, 0x35, 0x3f, 0xa2, 0xde, 0xac, 0x0b, 0x3c, 0xfd, 0x9f,
	0xfd, 0xd0, 0xc9, 0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x68, 0x29, 0xa2, 0x00, 0x60, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigurationClient is the client API for Configuration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigurationClient interface {
	SetConnection(ctx context.Context, in *Conn, opts ...grpc.CallOption) (*Result, error)
	SetAffinity(ctx context.Context, in *Affinity, opts ...grpc.CallOption) (*Result, error)
	GetGoroutineCount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Count, error)
	UpdateLogLevel(ctx context.Context, in *common.Logging, opts ...grpc.CallOption) (*Empty, error)
	GetLogLevels(ctx context.Context, in *common.LoggingComponent, opts ...grpc.CallOption) (*common.Loggings, error)
}

type configurationClient struct {
	cc *grpc.ClientConn
}

func NewConfigurationClient(cc *grpc.ClientConn) ConfigurationClient {
	return &configurationClient{cc}
}

func (c *configurationClient) SetConnection(ctx context.Context, in *Conn, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/afrouter.Configuration/SetConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) SetAffinity(ctx context.Context, in *Affinity, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/afrouter.Configuration/SetAffinity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) GetGoroutineCount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/afrouter.Configuration/GetGoroutineCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) UpdateLogLevel(ctx context.Context, in *common.Logging, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/afrouter.Configuration/UpdateLogLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) GetLogLevels(ctx context.Context, in *common.LoggingComponent, opts ...grpc.CallOption) (*common.Loggings, error) {
	out := new(common.Loggings)
	err := c.cc.Invoke(ctx, "/afrouter.Configuration/GetLogLevels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServer is the server API for Configuration service.
type ConfigurationServer interface {
	SetConnection(context.Context, *Conn) (*Result, error)
	SetAffinity(context.Context, *Affinity) (*Result, error)
	GetGoroutineCount(context.Context, *Empty) (*Count, error)
	UpdateLogLevel(context.Context, *common.Logging) (*Empty, error)
	GetLogLevels(context.Context, *common.LoggingComponent) (*common.Loggings, error)
}

func RegisterConfigurationServer(s *grpc.Server, srv ConfigurationServer) {
	s.RegisterService(&_Configuration_serviceDesc, srv)
}

func _Configuration_SetConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Conn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).SetConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/afrouter.Configuration/SetConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).SetConnection(ctx, req.(*Conn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_SetAffinity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Affinity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).SetAffinity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/afrouter.Configuration/SetAffinity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).SetAffinity(ctx, req.(*Affinity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_GetGoroutineCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).GetGoroutineCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/afrouter.Configuration/GetGoroutineCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).GetGoroutineCount(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_UpdateLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Logging)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).UpdateLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/afrouter.Configuration/UpdateLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).UpdateLogLevel(ctx, req.(*common.Logging))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_GetLogLevels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.LoggingComponent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).GetLogLevels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/afrouter.Configuration/GetLogLevels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).GetLogLevels(ctx, req.(*common.LoggingComponent))
	}
	return interceptor(ctx, in, info, handler)
}

var _Configuration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "afrouter.Configuration",
	HandlerType: (*ConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetConnection",
			Handler:    _Configuration_SetConnection_Handler,
		},
		{
			MethodName: "SetAffinity",
			Handler:    _Configuration_SetAffinity_Handler,
		},
		{
			MethodName: "GetGoroutineCount",
			Handler:    _Configuration_GetGoroutineCount_Handler,
		},
		{
			MethodName: "UpdateLogLevel",
			Handler:    _Configuration_UpdateLogLevel_Handler,
		},
		{
			MethodName: "GetLogLevels",
			Handler:    _Configuration_GetLogLevels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "voltha_protos/afrouter.proto",
}
