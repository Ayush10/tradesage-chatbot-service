// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: proto/chatbot.proto

package pb

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
	ChatbotService_SendMessage_FullMethodName = "/chatbot.ChatbotService/SendMessage"
)

// ChatbotServiceClient is the client API for ChatbotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The Chatbot service definition.
type ChatbotServiceClient interface {
	// Sends a message to the chatbot and receives a response.
	SendMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error)
}

type chatbotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatbotServiceClient(cc grpc.ClientConnInterface) ChatbotServiceClient {
	return &chatbotServiceClient{cc}
}

func (c *chatbotServiceClient) SendMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatResponse)
	err := c.cc.Invoke(ctx, ChatbotService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatbotServiceServer is the server API for ChatbotService service.
// All implementations must embed UnimplementedChatbotServiceServer
// for forward compatibility.
//
// The Chatbot service definition.
type ChatbotServiceServer interface {
	// Sends a message to the chatbot and receives a response.
	SendMessage(context.Context, *ChatRequest) (*ChatResponse, error)
	mustEmbedUnimplementedChatbotServiceServer()
}

// UnimplementedChatbotServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatbotServiceServer struct{}

func (UnimplementedChatbotServiceServer) SendMessage(context.Context, *ChatRequest) (*ChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatbotServiceServer) mustEmbedUnimplementedChatbotServiceServer() {}
func (UnimplementedChatbotServiceServer) testEmbeddedByValue()                        {}

// UnsafeChatbotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatbotServiceServer will
// result in compilation errors.
type UnsafeChatbotServiceServer interface {
	mustEmbedUnimplementedChatbotServiceServer()
}

func RegisterChatbotServiceServer(s grpc.ServiceRegistrar, srv ChatbotServiceServer) {
	// If the following call pancis, it indicates UnimplementedChatbotServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatbotService_ServiceDesc, srv)
}

func _ChatbotService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatbotServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatbotService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatbotServiceServer).SendMessage(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatbotService_ServiceDesc is the grpc.ServiceDesc for ChatbotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatbotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatbot.ChatbotService",
	HandlerType: (*ChatbotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ChatbotService_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/chatbot.proto",
}
