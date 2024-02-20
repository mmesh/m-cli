// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0--rc2
// source: mmesh/protobuf/rpc/v1/billingAPI.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	empty "mmesh.dev/m-api-go/grpc/common/empty"
	billing "mmesh.dev/m-api-go/grpc/resources/billing"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BillingAPI_SearchCustomer_FullMethodName    = "/api.BillingAPI/SearchCustomer"
	BillingAPI_GetCustomerPortal_FullMethodName = "/api.BillingAPI/GetCustomerPortal"
)

// BillingAPIClient is the client API for BillingAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillingAPIClient interface {
	// customer
	SearchCustomer(ctx context.Context, in *billing.Customer, opts ...grpc.CallOption) (*empty.Response, error)
	GetCustomerPortal(ctx context.Context, in *billing.CustomerPortalRequest, opts ...grpc.CallOption) (*billing.CustomerPortalResponse, error)
}

type billingAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingAPIClient(cc grpc.ClientConnInterface) BillingAPIClient {
	return &billingAPIClient{cc}
}

func (c *billingAPIClient) SearchCustomer(ctx context.Context, in *billing.Customer, opts ...grpc.CallOption) (*empty.Response, error) {
	out := new(empty.Response)
	err := c.cc.Invoke(ctx, BillingAPI_SearchCustomer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingAPIClient) GetCustomerPortal(ctx context.Context, in *billing.CustomerPortalRequest, opts ...grpc.CallOption) (*billing.CustomerPortalResponse, error) {
	out := new(billing.CustomerPortalResponse)
	err := c.cc.Invoke(ctx, BillingAPI_GetCustomerPortal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingAPIServer is the server API for BillingAPI service.
// All implementations must embed UnimplementedBillingAPIServer
// for forward compatibility
type BillingAPIServer interface {
	// customer
	SearchCustomer(context.Context, *billing.Customer) (*empty.Response, error)
	GetCustomerPortal(context.Context, *billing.CustomerPortalRequest) (*billing.CustomerPortalResponse, error)
	mustEmbedUnimplementedBillingAPIServer()
}

// UnimplementedBillingAPIServer must be embedded to have forward compatible implementations.
type UnimplementedBillingAPIServer struct {
}

func (UnimplementedBillingAPIServer) SearchCustomer(context.Context, *billing.Customer) (*empty.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCustomer not implemented")
}
func (UnimplementedBillingAPIServer) GetCustomerPortal(context.Context, *billing.CustomerPortalRequest) (*billing.CustomerPortalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerPortal not implemented")
}
func (UnimplementedBillingAPIServer) mustEmbedUnimplementedBillingAPIServer() {}

// UnsafeBillingAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillingAPIServer will
// result in compilation errors.
type UnsafeBillingAPIServer interface {
	mustEmbedUnimplementedBillingAPIServer()
}

func RegisterBillingAPIServer(s grpc.ServiceRegistrar, srv BillingAPIServer) {
	s.RegisterService(&BillingAPI_ServiceDesc, srv)
}

func _BillingAPI_SearchCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingAPIServer).SearchCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingAPI_SearchCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingAPIServer).SearchCustomer(ctx, req.(*billing.Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingAPI_GetCustomerPortal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(billing.CustomerPortalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingAPIServer).GetCustomerPortal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingAPI_GetCustomerPortal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingAPIServer).GetCustomerPortal(ctx, req.(*billing.CustomerPortalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BillingAPI_ServiceDesc is the grpc.ServiceDesc for BillingAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BillingAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.BillingAPI",
	HandlerType: (*BillingAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchCustomer",
			Handler:    _BillingAPI_SearchCustomer_Handler,
		},
		{
			MethodName: "GetCustomerPortal",
			Handler:    _BillingAPI_GetCustomerPortal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mmesh/protobuf/rpc/v1/billingAPI.proto",
}
