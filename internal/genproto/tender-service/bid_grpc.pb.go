// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: bid.proto

package tender_service

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

// BidServiceClient is the client API for BidService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BidServiceClient interface {
	SubmitBid(ctx context.Context, in *SubmitBidRequest, opts ...grpc.CallOption) (*BidResponse, error)
	GetBidsForTender(ctx context.Context, in *GetBidsRequest, opts ...grpc.CallOption) (*BidsListResponse, error)
}

type bidServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBidServiceClient(cc grpc.ClientConnInterface) BidServiceClient {
	return &bidServiceClient{cc}
}

func (c *bidServiceClient) SubmitBid(ctx context.Context, in *SubmitBidRequest, opts ...grpc.CallOption) (*BidResponse, error) {
	out := new(BidResponse)
	err := c.cc.Invoke(ctx, "/bid.BidService/SubmitBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bidServiceClient) GetBidsForTender(ctx context.Context, in *GetBidsRequest, opts ...grpc.CallOption) (*BidsListResponse, error) {
	out := new(BidsListResponse)
	err := c.cc.Invoke(ctx, "/bid.BidService/GetBidsForTender", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BidServiceServer is the server API for BidService service.
// All implementations must embed UnimplementedBidServiceServer
// for forward compatibility
type BidServiceServer interface {
	SubmitBid(context.Context, *SubmitBidRequest) (*BidResponse, error)
	GetBidsForTender(context.Context, *GetBidsRequest) (*BidsListResponse, error)
	mustEmbedUnimplementedBidServiceServer()
}

// UnimplementedBidServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBidServiceServer struct {
}

func (UnimplementedBidServiceServer) SubmitBid(context.Context, *SubmitBidRequest) (*BidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitBid not implemented")
}
func (UnimplementedBidServiceServer) GetBidsForTender(context.Context, *GetBidsRequest) (*BidsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBidsForTender not implemented")
}
func (UnimplementedBidServiceServer) mustEmbedUnimplementedBidServiceServer() {}

// UnsafeBidServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BidServiceServer will
// result in compilation errors.
type UnsafeBidServiceServer interface {
	mustEmbedUnimplementedBidServiceServer()
}

func RegisterBidServiceServer(s grpc.ServiceRegistrar, srv BidServiceServer) {
	s.RegisterService(&BidService_ServiceDesc, srv)
}

func _BidService_SubmitBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitBidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidServiceServer).SubmitBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bid.BidService/SubmitBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidServiceServer).SubmitBid(ctx, req.(*SubmitBidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BidService_GetBidsForTender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBidsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BidServiceServer).GetBidsForTender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bid.BidService/GetBidsForTender",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BidServiceServer).GetBidsForTender(ctx, req.(*GetBidsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BidService_ServiceDesc is the grpc.ServiceDesc for BidService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BidService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bid.BidService",
	HandlerType: (*BidServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitBid",
			Handler:    _BidService_SubmitBid_Handler,
		},
		{
			MethodName: "GetBidsForTender",
			Handler:    _BidService_GetBidsForTender_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bid.proto",
}