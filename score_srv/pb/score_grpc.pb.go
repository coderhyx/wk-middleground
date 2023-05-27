// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: score.proto

package __

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

const (
	Score_Create_FullMethodName      = "/pb.score/create"
	Score_Update_FullMethodName      = "/pb.score/update"
	Score_DtmCreate_FullMethodName   = "/pb.score/dtmCreate"
	Score_DtmRollback_FullMethodName = "/pb.score/dtmRollback"
	Score_GetScore_FullMethodName    = "/pb.score/getScore"
)

// ScoreClient is the client API for Score service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScoreClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*CreateResp, error)
	DtmCreate(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	DtmRollback(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	GetScore(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
}

type scoreClient struct {
	cc grpc.ClientConnInterface
}

func NewScoreClient(cc grpc.ClientConnInterface) ScoreClient {
	return &scoreClient{cc}
}

func (c *scoreClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, Score_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, Score_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreClient) DtmCreate(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, Score_DtmCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreClient) DtmRollback(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, Score_DtmRollback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreClient) GetScore(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, Score_GetScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScoreServer is the server API for Score service.
// All implementations must embed UnimplementedScoreServer
// for forward compatibility
type ScoreServer interface {
	Create(context.Context, *CreateReq) (*CreateResp, error)
	Update(context.Context, *UpdateReq) (*CreateResp, error)
	DtmCreate(context.Context, *CreateReq) (*CreateResp, error)
	DtmRollback(context.Context, *CreateReq) (*CreateResp, error)
	GetScore(context.Context, *GetReq) (*GetResp, error)
	mustEmbedUnimplementedScoreServer()
}

// UnimplementedScoreServer must be embedded to have forward compatible implementations.
type UnimplementedScoreServer struct {
}

func (UnimplementedScoreServer) Create(context.Context, *CreateReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedScoreServer) Update(context.Context, *UpdateReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedScoreServer) DtmCreate(context.Context, *CreateReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DtmCreate not implemented")
}
func (UnimplementedScoreServer) DtmRollback(context.Context, *CreateReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DtmRollback not implemented")
}
func (UnimplementedScoreServer) GetScore(context.Context, *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScore not implemented")
}
func (UnimplementedScoreServer) mustEmbedUnimplementedScoreServer() {}

// UnsafeScoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScoreServer will
// result in compilation errors.
type UnsafeScoreServer interface {
	mustEmbedUnimplementedScoreServer()
}

func RegisterScoreServer(s grpc.ServiceRegistrar, srv ScoreServer) {
	s.RegisterService(&Score_ServiceDesc, srv)
}

func _Score_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Score_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Score_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Score_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Score_DtmCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServer).DtmCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Score_DtmCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServer).DtmCreate(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Score_DtmRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServer).DtmRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Score_DtmRollback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServer).DtmRollback(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Score_GetScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServer).GetScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Score_GetScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServer).GetScore(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Score_ServiceDesc is the grpc.ServiceDesc for Score service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Score_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.score",
	HandlerType: (*ScoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Score_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Score_Update_Handler,
		},
		{
			MethodName: "dtmCreate",
			Handler:    _Score_DtmCreate_Handler,
		},
		{
			MethodName: "dtmRollback",
			Handler:    _Score_DtmRollback_Handler,
		},
		{
			MethodName: "getScore",
			Handler:    _Score_GetScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "score.proto",
}