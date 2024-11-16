// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: api/proto/gnss-radar/gnss-radar.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GnssService_GetStatus_FullMethodName = "/gnss_radar.GnssService/GetStatus"
)

// GnssServiceClient is the client API for GnssService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GnssServiceClient interface {
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
}

type gnssServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGnssServiceClient(cc grpc.ClientConnInterface) GnssServiceClient {
	return &gnssServiceClient{cc}
}

func (c *gnssServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, GnssService_GetStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GnssServiceServer is the server API for GnssService service.
// All implementations must embed UnimplementedGnssServiceServer
// for forward compatibility.
type GnssServiceServer interface {
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	mustEmbedUnimplementedGnssServiceServer()
}

// UnimplementedGnssServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGnssServiceServer struct{}

func (UnimplementedGnssServiceServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedGnssServiceServer) mustEmbedUnimplementedGnssServiceServer() {}
func (UnimplementedGnssServiceServer) testEmbeddedByValue()                     {}

// UnsafeGnssServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GnssServiceServer will
// result in compilation errors.
type UnsafeGnssServiceServer interface {
	mustEmbedUnimplementedGnssServiceServer()
}

func RegisterGnssServiceServer(s grpc.ServiceRegistrar, srv GnssServiceServer) {
	// If the following call pancis, it indicates UnimplementedGnssServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GnssService_ServiceDesc, srv)
}

func _GnssService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnssServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GnssService_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnssServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GnssService_ServiceDesc is the grpc.ServiceDesc for GnssService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GnssService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnss_radar.GnssService",
	HandlerType: (*GnssServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _GnssService_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/gnss-radar/gnss-radar.proto",
}