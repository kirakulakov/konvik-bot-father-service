// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: bot_father.proto

package bot_father_grpc

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

// BotFatherClient is the client API for BotFather service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotFatherClient interface {
	GetBot(ctx context.Context, in *GetBotRequest, opts ...grpc.CallOption) (*GetBotResponse, error)
}

type botFatherClient struct {
	cc grpc.ClientConnInterface
}

func NewBotFatherClient(cc grpc.ClientConnInterface) BotFatherClient {
	return &botFatherClient{cc}
}

func (c *botFatherClient) GetBot(ctx context.Context, in *GetBotRequest, opts ...grpc.CallOption) (*GetBotResponse, error) {
	out := new(GetBotResponse)
	err := c.cc.Invoke(ctx, "/bot_father.BotFather/GetBot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotFatherServer is the server API for BotFather service.
// All implementations must embed UnimplementedBotFatherServer
// for forward compatibility
type BotFatherServer interface {
	GetBot(context.Context, *GetBotRequest) (*GetBotResponse, error)
	mustEmbedUnimplementedBotFatherServer()
}

// UnimplementedBotFatherServer must be embedded to have forward compatible implementations.
type UnimplementedBotFatherServer struct {
}

func (UnimplementedBotFatherServer) GetBot(context.Context, *GetBotRequest) (*GetBotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBot not implemented")
}
func (UnimplementedBotFatherServer) mustEmbedUnimplementedBotFatherServer() {}

// UnsafeBotFatherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotFatherServer will
// result in compilation errors.
type UnsafeBotFatherServer interface {
	mustEmbedUnimplementedBotFatherServer()
}

func RegisterBotFatherServer(s grpc.ServiceRegistrar, srv BotFatherServer) {
	s.RegisterService(&BotFather_ServiceDesc, srv)
}

func _BotFather_GetBot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotFatherServer).GetBot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bot_father.BotFather/GetBot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotFatherServer).GetBot(ctx, req.(*GetBotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BotFather_ServiceDesc is the grpc.ServiceDesc for BotFather service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BotFather_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bot_father.BotFather",
	HandlerType: (*BotFatherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBot",
			Handler:    _BotFather_GetBot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bot_father.proto",
}