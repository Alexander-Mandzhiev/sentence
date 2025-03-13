// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: departments/proto.proto

package departments

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
	DepartmentsProvider_Create_FullMethodName = "/departments_provider.DepartmentsProvider/Create"
	DepartmentsProvider_Get_FullMethodName    = "/departments_provider.DepartmentsProvider/Get"
	DepartmentsProvider_Update_FullMethodName = "/departments_provider.DepartmentsProvider/Update"
	DepartmentsProvider_Delete_FullMethodName = "/departments_provider.DepartmentsProvider/Delete"
	DepartmentsProvider_List_FullMethodName   = "/departments_provider.DepartmentsProvider/List"
)

// DepartmentsProviderClient is the client API for DepartmentsProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepartmentsProviderClient interface {
	Create(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*DepartmentResponse, error)
	Get(ctx context.Context, in *GetDepartmentRequest, opts ...grpc.CallOption) (*DepartmentResponse, error)
	Update(ctx context.Context, in *UpdateDepartmentRequest, opts ...grpc.CallOption) (*DepartmentResponse, error)
	Delete(ctx context.Context, in *DeleteDepartmentRequest, opts ...grpc.CallOption) (*DeleteDepartmentResponse, error)
	List(ctx context.Context, in *ListDepartmentsRequest, opts ...grpc.CallOption) (*DepartmentsListResponse, error)
}

type departmentsProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewDepartmentsProviderClient(cc grpc.ClientConnInterface) DepartmentsProviderClient {
	return &departmentsProviderClient{cc}
}

func (c *departmentsProviderClient) Create(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*DepartmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentsProvider_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsProviderClient) Get(ctx context.Context, in *GetDepartmentRequest, opts ...grpc.CallOption) (*DepartmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentsProvider_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsProviderClient) Update(ctx context.Context, in *UpdateDepartmentRequest, opts ...grpc.CallOption) (*DepartmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentsProvider_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsProviderClient) Delete(ctx context.Context, in *DeleteDepartmentRequest, opts ...grpc.CallOption) (*DeleteDepartmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentsProvider_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentsProviderClient) List(ctx context.Context, in *ListDepartmentsRequest, opts ...grpc.CallOption) (*DepartmentsListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DepartmentsListResponse)
	err := c.cc.Invoke(ctx, DepartmentsProvider_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepartmentsProviderServer is the server API for DepartmentsProvider service.
// All implementations must embed UnimplementedDepartmentsProviderServer
// for forward compatibility.
type DepartmentsProviderServer interface {
	Create(context.Context, *CreateDepartmentRequest) (*DepartmentResponse, error)
	Get(context.Context, *GetDepartmentRequest) (*DepartmentResponse, error)
	Update(context.Context, *UpdateDepartmentRequest) (*DepartmentResponse, error)
	Delete(context.Context, *DeleteDepartmentRequest) (*DeleteDepartmentResponse, error)
	List(context.Context, *ListDepartmentsRequest) (*DepartmentsListResponse, error)
	mustEmbedUnimplementedDepartmentsProviderServer()
}

// UnimplementedDepartmentsProviderServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDepartmentsProviderServer struct{}

func (UnimplementedDepartmentsProviderServer) Create(context.Context, *CreateDepartmentRequest) (*DepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDepartmentsProviderServer) Get(context.Context, *GetDepartmentRequest) (*DepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDepartmentsProviderServer) Update(context.Context, *UpdateDepartmentRequest) (*DepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDepartmentsProviderServer) Delete(context.Context, *DeleteDepartmentRequest) (*DeleteDepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDepartmentsProviderServer) List(context.Context, *ListDepartmentsRequest) (*DepartmentsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDepartmentsProviderServer) mustEmbedUnimplementedDepartmentsProviderServer() {}
func (UnimplementedDepartmentsProviderServer) testEmbeddedByValue()                             {}

// UnsafeDepartmentsProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepartmentsProviderServer will
// result in compilation errors.
type UnsafeDepartmentsProviderServer interface {
	mustEmbedUnimplementedDepartmentsProviderServer()
}

func RegisterDepartmentsProviderServer(s grpc.ServiceRegistrar, srv DepartmentsProviderServer) {
	// If the following call pancis, it indicates UnimplementedDepartmentsProviderServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DepartmentsProvider_ServiceDesc, srv)
}

func _DepartmentsProvider_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentsProviderServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentsProvider_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentsProviderServer).Create(ctx, req.(*CreateDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentsProvider_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentsProviderServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentsProvider_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentsProviderServer).Get(ctx, req.(*GetDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentsProvider_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentsProviderServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentsProvider_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentsProviderServer).Update(ctx, req.(*UpdateDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentsProvider_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentsProviderServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentsProvider_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentsProviderServer).Delete(ctx, req.(*DeleteDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentsProvider_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDepartmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentsProviderServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentsProvider_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentsProviderServer).List(ctx, req.(*ListDepartmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DepartmentsProvider_ServiceDesc is the grpc.ServiceDesc for DepartmentsProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepartmentsProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "departments_provider.DepartmentsProvider",
	HandlerType: (*DepartmentsProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DepartmentsProvider_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DepartmentsProvider_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DepartmentsProvider_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DepartmentsProvider_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DepartmentsProvider_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "departments/proto.proto",
}
