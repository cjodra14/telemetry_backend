// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: telemetry.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TelemetryServiceClient is the client API for TelemetryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TelemetryServiceClient interface {
	SaveTelemetry(ctx context.Context, in *Telemetry, opts ...grpc.CallOption) (*TelemetryResponse, error)
	GetUserTelemetry(ctx context.Context, in *UserTelemetryRequest, opts ...grpc.CallOption) (*UserTelemetryResponse, error)
}

type telemetryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTelemetryServiceClient(cc grpc.ClientConnInterface) TelemetryServiceClient {
	return &telemetryServiceClient{cc}
}

func (c *telemetryServiceClient) SaveTelemetry(ctx context.Context, in *Telemetry, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/TelemetryService/SaveTelemetry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServiceClient) GetUserTelemetry(ctx context.Context, in *UserTelemetryRequest, opts ...grpc.CallOption) (*UserTelemetryResponse, error) {
	out := new(UserTelemetryResponse)
	err := c.cc.Invoke(ctx, "/TelemetryService/GetUserTelemetry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelemetryServiceServer is the server API for TelemetryService service.
// All implementations must embed UnimplementedTelemetryServiceServer
// for forward compatibility
type TelemetryServiceServer interface {
	SaveTelemetry(context.Context, *Telemetry) (*TelemetryResponse, error)
	GetUserTelemetry(context.Context, *UserTelemetryRequest) (*UserTelemetryResponse, error)
	mustEmbedUnimplementedTelemetryServiceServer()
}

// UnimplementedTelemetryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTelemetryServiceServer struct {
}

func (UnimplementedTelemetryServiceServer) SaveTelemetry(context.Context, *Telemetry) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTelemetry not implemented")
}
func (UnimplementedTelemetryServiceServer) GetUserTelemetry(context.Context, *UserTelemetryRequest) (*UserTelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTelemetry not implemented")
}
func (UnimplementedTelemetryServiceServer) mustEmbedUnimplementedTelemetryServiceServer() {}

// UnsafeTelemetryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelemetryServiceServer will
// result in compilation errors.
type UnsafeTelemetryServiceServer interface {
	mustEmbedUnimplementedTelemetryServiceServer()
}

func RegisterTelemetryServiceServer(s grpc.ServiceRegistrar, srv TelemetryServiceServer) {
	s.RegisterService(&TelemetryService_ServiceDesc, srv)
}

func _TelemetryService_SaveTelemetry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Telemetry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServiceServer).SaveTelemetry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TelemetryService/SaveTelemetry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServiceServer).SaveTelemetry(ctx, req.(*Telemetry))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryService_GetUserTelemetry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServiceServer).GetUserTelemetry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TelemetryService/GetUserTelemetry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServiceServer).GetUserTelemetry(ctx, req.(*UserTelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TelemetryService_ServiceDesc is the grpc.ServiceDesc for TelemetryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TelemetryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TelemetryService",
	HandlerType: (*TelemetryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveTelemetry",
			Handler:    _TelemetryService_SaveTelemetry_Handler,
		},
		{
			MethodName: "GetUserTelemetry",
			Handler:    _TelemetryService_GetUserTelemetry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "telemetry.proto",
}