// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/proto/boggle.proto

package api

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

// BoggleServiceClient is the client API for BoggleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoggleServiceClient interface {
	Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error)
	Solution(ctx context.Context, in *SolutionRequest, opts ...grpc.CallOption) (*SolutionResponse, error)
}

type boggleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBoggleServiceClient(cc grpc.ClientConnInterface) BoggleServiceClient {
	return &boggleServiceClient{cc}
}

func (c *boggleServiceClient) Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error) {
	out := new(SolveResponse)
	err := c.cc.Invoke(ctx, "/api.BoggleService/Solve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boggleServiceClient) Solution(ctx context.Context, in *SolutionRequest, opts ...grpc.CallOption) (*SolutionResponse, error) {
	out := new(SolutionResponse)
	err := c.cc.Invoke(ctx, "/api.BoggleService/Solution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoggleServiceServer is the server API for BoggleService service.
// All implementations must embed UnimplementedBoggleServiceServer
// for forward compatibility
type BoggleServiceServer interface {
	Solve(context.Context, *SolveRequest) (*SolveResponse, error)
	Solution(context.Context, *SolutionRequest) (*SolutionResponse, error)
	mustEmbedUnimplementedBoggleServiceServer()
}

// UnimplementedBoggleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBoggleServiceServer struct {
}

func (UnimplementedBoggleServiceServer) Solve(context.Context, *SolveRequest) (*SolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Solve not implemented")
}
func (UnimplementedBoggleServiceServer) Solution(context.Context, *SolutionRequest) (*SolutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Solution not implemented")
}
func (UnimplementedBoggleServiceServer) mustEmbedUnimplementedBoggleServiceServer() {}

// UnsafeBoggleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoggleServiceServer will
// result in compilation errors.
type UnsafeBoggleServiceServer interface {
	mustEmbedUnimplementedBoggleServiceServer()
}

func RegisterBoggleServiceServer(s grpc.ServiceRegistrar, srv BoggleServiceServer) {
	s.RegisterService(&BoggleService_ServiceDesc, srv)
}

func _BoggleService_Solve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoggleServiceServer).Solve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BoggleService/Solve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoggleServiceServer).Solve(ctx, req.(*SolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoggleService_Solution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoggleServiceServer).Solution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BoggleService/Solution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoggleServiceServer).Solution(ctx, req.(*SolutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BoggleService_ServiceDesc is the grpc.ServiceDesc for BoggleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BoggleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.BoggleService",
	HandlerType: (*BoggleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Solve",
			Handler:    _BoggleService_Solve_Handler,
		},
		{
			MethodName: "Solution",
			Handler:    _BoggleService_Solution_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/boggle.proto",
}