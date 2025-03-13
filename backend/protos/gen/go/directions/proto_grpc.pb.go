// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: directions/proto.proto

package directions

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
	DirectionsProvider_Create_FullMethodName = "/directions_provider.DirectionsProvider/Create"
	DirectionsProvider_Get_FullMethodName    = "/directions_provider.DirectionsProvider/Get"
	DirectionsProvider_Update_FullMethodName = "/directions_provider.DirectionsProvider/Update"
	DirectionsProvider_Delete_FullMethodName = "/directions_provider.DirectionsProvider/Delete"
	DirectionsProvider_List_FullMethodName   = "/directions_provider.DirectionsProvider/List"
)

// DirectionsProviderClient is the client API for DirectionsProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectionsProviderClient interface {
	Create(ctx context.Context, in *CreateDirectionRequest, opts ...grpc.CallOption) (*DirectionResponse, error)
	Get(ctx context.Context, in *GetDirectionRequest, opts ...grpc.CallOption) (*DirectionResponse, error)
	Update(ctx context.Context, in *UpdateDirectionRequest, opts ...grpc.CallOption) (*DirectionResponse, error)
	Delete(ctx context.Context, in *DeleteDirectionRequest, opts ...grpc.CallOption) (*DeleteDirectionResponse, error)
	List(ctx context.Context, in *ListDirectionsRequest, opts ...grpc.CallOption) (*DirectionsListResponse, error)
}

type directionsProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectionsProviderClient(cc grpc.ClientConnInterface) DirectionsProviderClient {
	return &directionsProviderClient{cc}
}

func (c *directionsProviderClient) Create(ctx context.Context, in *CreateDirectionRequest, opts ...grpc.CallOption) (*DirectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DirectionResponse)
	err := c.cc.Invoke(ctx, DirectionsProvider_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directionsProviderClient) Get(ctx context.Context, in *GetDirectionRequest, opts ...grpc.CallOption) (*DirectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DirectionResponse)
	err := c.cc.Invoke(ctx, DirectionsProvider_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directionsProviderClient) Update(ctx context.Context, in *UpdateDirectionRequest, opts ...grpc.CallOption) (*DirectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DirectionResponse)
	err := c.cc.Invoke(ctx, DirectionsProvider_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directionsProviderClient) Delete(ctx context.Context, in *DeleteDirectionRequest, opts ...grpc.CallOption) (*DeleteDirectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDirectionResponse)
	err := c.cc.Invoke(ctx, DirectionsProvider_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directionsProviderClient) List(ctx context.Context, in *ListDirectionsRequest, opts ...grpc.CallOption) (*DirectionsListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DirectionsListResponse)
	err := c.cc.Invoke(ctx, DirectionsProvider_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectionsProviderServer is the server API for DirectionsProvider service.
// All implementations must embed UnimplementedDirectionsProviderServer
// for forward compatibility.
type DirectionsProviderServer interface {
	Create(context.Context, *CreateDirectionRequest) (*DirectionResponse, error)
	Get(context.Context, *GetDirectionRequest) (*DirectionResponse, error)
	Update(context.Context, *UpdateDirectionRequest) (*DirectionResponse, error)
	Delete(context.Context, *DeleteDirectionRequest) (*DeleteDirectionResponse, error)
	List(context.Context, *ListDirectionsRequest) (*DirectionsListResponse, error)
	mustEmbedUnimplementedDirectionsProviderServer()
}

// UnimplementedDirectionsProviderServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDirectionsProviderServer struct{}

func (UnimplementedDirectionsProviderServer) Create(context.Context, *CreateDirectionRequest) (*DirectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDirectionsProviderServer) Get(context.Context, *GetDirectionRequest) (*DirectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDirectionsProviderServer) Update(context.Context, *UpdateDirectionRequest) (*DirectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDirectionsProviderServer) Delete(context.Context, *DeleteDirectionRequest) (*DeleteDirectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDirectionsProviderServer) List(context.Context, *ListDirectionsRequest) (*DirectionsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDirectionsProviderServer) mustEmbedUnimplementedDirectionsProviderServer() {}
func (UnimplementedDirectionsProviderServer) testEmbeddedByValue()                            {}

// UnsafeDirectionsProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectionsProviderServer will
// result in compilation errors.
type UnsafeDirectionsProviderServer interface {
	mustEmbedUnimplementedDirectionsProviderServer()
}

func RegisterDirectionsProviderServer(s grpc.ServiceRegistrar, srv DirectionsProviderServer) {
	// If the following call pancis, it indicates UnimplementedDirectionsProviderServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DirectionsProvider_ServiceDesc, srv)
}

func _DirectionsProvider_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDirectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectionsProviderServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DirectionsProvider_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectionsProviderServer).Create(ctx, req.(*CreateDirectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectionsProvider_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDirectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectionsProviderServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DirectionsProvider_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectionsProviderServer).Get(ctx, req.(*GetDirectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectionsProvider_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDirectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectionsProviderServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DirectionsProvider_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectionsProviderServer).Update(ctx, req.(*UpdateDirectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectionsProvider_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDirectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectionsProviderServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DirectionsProvider_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectionsProviderServer).Delete(ctx, req.(*DeleteDirectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectionsProvider_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDirectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectionsProviderServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DirectionsProvider_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectionsProviderServer).List(ctx, req.(*ListDirectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DirectionsProvider_ServiceDesc is the grpc.ServiceDesc for DirectionsProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DirectionsProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "directions_provider.DirectionsProvider",
	HandlerType: (*DirectionsProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DirectionsProvider_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DirectionsProvider_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DirectionsProvider_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DirectionsProvider_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DirectionsProvider_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "directions/proto.proto",
}
