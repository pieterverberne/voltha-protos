// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/onu_inter_adapter_service.proto

package onu_inter_adapter_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	health "github.com/opencord/voltha-protos/v5/go/health"
	inter_adapter "github.com/opencord/voltha-protos/v5/go/inter_adapter"
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

func init() {
	proto.RegisterFile("voltha_protos/onu_inter_adapter_service.proto", fileDescriptor_f951f30caeee9ccd)
}

var fileDescriptor_f951f30caeee9ccd = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0xb9, 0x5c, 0xb8, 0x8b, 0xb9, 0x54, 0x61, 0x94, 0x82, 0xd1, 0x4d, 0x71, 0xe5, 0xa2,
	0x13, 0x50, 0x5c, 0x4b, 0x6b, 0xa5, 0x15, 0x2a, 0x2d, 0xb4, 0x88, 0xb8, 0x29, 0xd3, 0xc9, 0x69,
	0x32, 0x90, 0xcc, 0x09, 0x93, 0x93, 0x8a, 0x6f, 0xe1, 0x1b, 0xfa, 0x2a, 0xd2, 0x4c, 0x0b, 0x9d,
	0xda, 0x90, 0xd5, 0x2c, 0xce, 0x77, 0xbe, 0xf3, 0x0f, 0xfc, 0xac, 0xbb, 0xc6, 0x94, 0x12, 0xb9,
	0xc8, 0x2d, 0x12, 0x16, 0x21, 0x9a, 0x72, 0xa1, 0x0d, 0x81, 0x5d, 0xc8, 0x48, 0xe6, 0x9b, 0xb7,
	0x00, 0xbb, 0xd6, 0x0a, 0x44, 0x05, 0xf0, 0x8b, 0x5a, 0x20, 0xb8, 0x8c, 0x11, 0xe3, 0x14, 0xc2,
	0x0a, 0x5c, 0x96, 0xab, 0x10, 0xb2, 0x9c, 0x3e, 0xdd, 0x5e, 0xd0, 0xf1, 0xcf, 0x78, 0x86, 0x2d,
	0x12, 0xf8, 0x48, 0x02, 0x32, 0xa5, 0xc4, 0xcd, 0x6e, 0xbf, 0xff, 0xb2, 0xf6, 0xc4, 0x94, 0xcf,
	0x9b, 0xb5, 0x9e, 0xdb, 0x9a, 0xb9, 0xb3, 0xfc, 0x81, 0x9d, 0x0e, 0x81, 0x46, 0x15, 0x3d, 0x23,
	0x49, 0x65, 0xc1, 0xdb, 0xc2, 0x45, 0x11, 0xbb, 0x28, 0xe2, 0x69, 0x13, 0x25, 0x38, 0x17, 0x5b,
	0xa9, 0x47, 0x8f, 0x59, 0xab, 0x52, 0x47, 0x5a, 0x49, 0xd2, 0x68, 0xf8, 0xb5, 0xf0, 0xe3, 0x79,
	0xd3, 0x17, 0x28, 0x0a, 0x19, 0x43, 0x50, 0x73, 0x83, 0x0f, 0xd8, 0xc9, 0x24, 0x53, 0x7a, 0x4f,
	0x17, 0x1c, 0xea, 0x32, 0xa5, 0x9b, 0x2c, 0x6f, 0xec, 0x6c, 0x80, 0x1f, 0x26, 0x45, 0x19, 0xcd,
	0x41, 0x25, 0x53, 0x8b, 0x2b, 0x9d, 0x02, 0xbf, 0x39, 0x50, 0xed, 0xcd, 0x76, 0x78, 0x93, 0x79,
	0xcc, 0x5a, 0x03, 0x48, 0x81, 0x60, 0x08, 0xd9, 0x14, 0x2d, 0xfd, 0xfa, 0xad, 0x37, 0x6d, 0xb2,
	0x8d, 0xd8, 0x7f, 0xc7, 0xcf, 0x1f, 0xd1, 0x10, 0xef, 0x1c, 0x75, 0xcd, 0x15, 0x9a, 0x26, 0x53,
	0xff, 0xeb, 0x0f, 0xeb, 0xa2, 0x8d, 0x05, 0xe6, 0x60, 0x14, 0xda, 0x48, 0xb8, 0x36, 0x88, 0xda,
	0xbe, 0xf5, 0xaf, 0x5e, 0x2b, 0xe2, 0x78, 0x2d, 0xde, 0x7b, 0xb1, 0xa6, 0xa4, 0x5c, 0x0a, 0x85,
	0x59, 0xb8, 0x73, 0x86, 0xce, 0xd9, 0xdd, 0x36, 0x6c, 0x7d, 0x1f, 0xc6, 0x58, 0xdf, 0xf8, 0xe5,
	0xbf, 0x8a, 0xbb, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x75, 0x2b, 0xd4, 0xd6, 0x23, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OnuInterAdapterServiceClient is the client API for OnuInterAdapterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnuInterAdapterServiceClient interface {
	// GetHealthStatus is used by an OnuInterAdapterService client to verify connectivity
	// to the gRPC server hosting the OnuInterAdapterService service
	GetHealthStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*health.HealthStatus, error)
	OnuIndication(ctx context.Context, in *inter_adapter.OnuIndicationMessage, opts ...grpc.CallOption) (*empty.Empty, error)
	OmciIndication(ctx context.Context, in *inter_adapter.OmciMessage, opts ...grpc.CallOption) (*empty.Empty, error)
	DownloadTechProfile(ctx context.Context, in *inter_adapter.TechProfileDownloadMessage, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteGemPort(ctx context.Context, in *inter_adapter.DeleteGemPortMessage, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteTCont(ctx context.Context, in *inter_adapter.DeleteTcontMessage, opts ...grpc.CallOption) (*empty.Empty, error)
}

type onuInterAdapterServiceClient struct {
	cc *grpc.ClientConn
}

func NewOnuInterAdapterServiceClient(cc *grpc.ClientConn) OnuInterAdapterServiceClient {
	return &onuInterAdapterServiceClient{cc}
}

func (c *onuInterAdapterServiceClient) GetHealthStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*health.HealthStatus, error) {
	out := new(health.HealthStatus)
	err := c.cc.Invoke(ctx, "/onu_inter_adapter_service.OnuInterAdapterService/GetHealthStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onuInterAdapterServiceClient) OnuIndication(ctx context.Context, in *inter_adapter.OnuIndicationMessage, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/onu_inter_adapter_service.OnuInterAdapterService/OnuIndication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onuInterAdapterServiceClient) OmciIndication(ctx context.Context, in *inter_adapter.OmciMessage, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/onu_inter_adapter_service.OnuInterAdapterService/OmciIndication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onuInterAdapterServiceClient) DownloadTechProfile(ctx context.Context, in *inter_adapter.TechProfileDownloadMessage, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/onu_inter_adapter_service.OnuInterAdapterService/DownloadTechProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onuInterAdapterServiceClient) DeleteGemPort(ctx context.Context, in *inter_adapter.DeleteGemPortMessage, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/onu_inter_adapter_service.OnuInterAdapterService/DeleteGemPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onuInterAdapterServiceClient) DeleteTCont(ctx context.Context, in *inter_adapter.DeleteTcontMessage, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/onu_inter_adapter_service.OnuInterAdapterService/DeleteTCont", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnuInterAdapterServiceServer is the server API for OnuInterAdapterService service.
type OnuInterAdapterServiceServer interface {
	// GetHealthStatus is used by an OnuInterAdapterService client to verify connectivity
	// to the gRPC server hosting the OnuInterAdapterService service
	GetHealthStatus(context.Context, *empty.Empty) (*health.HealthStatus, error)
	OnuIndication(context.Context, *inter_adapter.OnuIndicationMessage) (*empty.Empty, error)
	OmciIndication(context.Context, *inter_adapter.OmciMessage) (*empty.Empty, error)
	DownloadTechProfile(context.Context, *inter_adapter.TechProfileDownloadMessage) (*empty.Empty, error)
	DeleteGemPort(context.Context, *inter_adapter.DeleteGemPortMessage) (*empty.Empty, error)
	DeleteTCont(context.Context, *inter_adapter.DeleteTcontMessage) (*empty.Empty, error)
}

// UnimplementedOnuInterAdapterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOnuInterAdapterServiceServer struct {
}

func (*UnimplementedOnuInterAdapterServiceServer) GetHealthStatus(ctx context.Context, req *empty.Empty) (*health.HealthStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealthStatus not implemented")
}
func (*UnimplementedOnuInterAdapterServiceServer) OnuIndication(ctx context.Context, req *inter_adapter.OnuIndicationMessage) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnuIndication not implemented")
}
func (*UnimplementedOnuInterAdapterServiceServer) OmciIndication(ctx context.Context, req *inter_adapter.OmciMessage) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OmciIndication not implemented")
}
func (*UnimplementedOnuInterAdapterServiceServer) DownloadTechProfile(ctx context.Context, req *inter_adapter.TechProfileDownloadMessage) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadTechProfile not implemented")
}
func (*UnimplementedOnuInterAdapterServiceServer) DeleteGemPort(ctx context.Context, req *inter_adapter.DeleteGemPortMessage) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGemPort not implemented")
}
func (*UnimplementedOnuInterAdapterServiceServer) DeleteTCont(ctx context.Context, req *inter_adapter.DeleteTcontMessage) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTCont not implemented")
}

func RegisterOnuInterAdapterServiceServer(s *grpc.Server, srv OnuInterAdapterServiceServer) {
	s.RegisterService(&_OnuInterAdapterService_serviceDesc, srv)
}

func _OnuInterAdapterService_GetHealthStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnuInterAdapterServiceServer).GetHealthStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onu_inter_adapter_service.OnuInterAdapterService/GetHealthStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnuInterAdapterServiceServer).GetHealthStatus(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnuInterAdapterService_OnuIndication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(inter_adapter.OnuIndicationMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnuInterAdapterServiceServer).OnuIndication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onu_inter_adapter_service.OnuInterAdapterService/OnuIndication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnuInterAdapterServiceServer).OnuIndication(ctx, req.(*inter_adapter.OnuIndicationMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnuInterAdapterService_OmciIndication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(inter_adapter.OmciMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnuInterAdapterServiceServer).OmciIndication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onu_inter_adapter_service.OnuInterAdapterService/OmciIndication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnuInterAdapterServiceServer).OmciIndication(ctx, req.(*inter_adapter.OmciMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnuInterAdapterService_DownloadTechProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(inter_adapter.TechProfileDownloadMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnuInterAdapterServiceServer).DownloadTechProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onu_inter_adapter_service.OnuInterAdapterService/DownloadTechProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnuInterAdapterServiceServer).DownloadTechProfile(ctx, req.(*inter_adapter.TechProfileDownloadMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnuInterAdapterService_DeleteGemPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(inter_adapter.DeleteGemPortMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnuInterAdapterServiceServer).DeleteGemPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onu_inter_adapter_service.OnuInterAdapterService/DeleteGemPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnuInterAdapterServiceServer).DeleteGemPort(ctx, req.(*inter_adapter.DeleteGemPortMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnuInterAdapterService_DeleteTCont_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(inter_adapter.DeleteTcontMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnuInterAdapterServiceServer).DeleteTCont(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onu_inter_adapter_service.OnuInterAdapterService/DeleteTCont",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnuInterAdapterServiceServer).DeleteTCont(ctx, req.(*inter_adapter.DeleteTcontMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnuInterAdapterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onu_inter_adapter_service.OnuInterAdapterService",
	HandlerType: (*OnuInterAdapterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealthStatus",
			Handler:    _OnuInterAdapterService_GetHealthStatus_Handler,
		},
		{
			MethodName: "OnuIndication",
			Handler:    _OnuInterAdapterService_OnuIndication_Handler,
		},
		{
			MethodName: "OmciIndication",
			Handler:    _OnuInterAdapterService_OmciIndication_Handler,
		},
		{
			MethodName: "DownloadTechProfile",
			Handler:    _OnuInterAdapterService_DownloadTechProfile_Handler,
		},
		{
			MethodName: "DeleteGemPort",
			Handler:    _OnuInterAdapterService_DeleteGemPort_Handler,
		},
		{
			MethodName: "DeleteTCont",
			Handler:    _OnuInterAdapterService_DeleteTCont_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "voltha_protos/onu_inter_adapter_service.proto",
}
