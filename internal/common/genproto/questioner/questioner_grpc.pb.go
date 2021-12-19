// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package questioner

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

// QuestionerServiceClient is the client API for QuestionerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionerServiceClient interface {
	GetAllQuestions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Questions, error)
	CalculateAllQuestions(ctx context.Context, in *Answers, opts ...grpc.CallOption) (*Staticstic, error)
}

type questionerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionerServiceClient(cc grpc.ClientConnInterface) QuestionerServiceClient {
	return &questionerServiceClient{cc}
}

func (c *questionerServiceClient) GetAllQuestions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Questions, error) {
	out := new(Questions)
	err := c.cc.Invoke(ctx, "/questioner.QuestionerService/GetAllQuestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionerServiceClient) CalculateAllQuestions(ctx context.Context, in *Answers, opts ...grpc.CallOption) (*Staticstic, error) {
	out := new(Staticstic)
	err := c.cc.Invoke(ctx, "/questioner.QuestionerService/CalculateAllQuestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionerServiceServer is the server API for QuestionerService service.
// All implementations should embed UnimplementedQuestionerServiceServer
// for forward compatibility
type QuestionerServiceServer interface {
	GetAllQuestions(context.Context, *Empty) (*Questions, error)
	CalculateAllQuestions(context.Context, *Answers) (*Staticstic, error)
}

// UnimplementedQuestionerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedQuestionerServiceServer struct {
}

func (UnimplementedQuestionerServiceServer) GetAllQuestions(context.Context, *Empty) (*Questions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllQuestions not implemented")
}
func (UnimplementedQuestionerServiceServer) CalculateAllQuestions(context.Context, *Answers) (*Staticstic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateAllQuestions not implemented")
}

// UnsafeQuestionerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionerServiceServer will
// result in compilation errors.
type UnsafeQuestionerServiceServer interface {
	mustEmbedUnimplementedQuestionerServiceServer()
}

func RegisterQuestionerServiceServer(s grpc.ServiceRegistrar, srv QuestionerServiceServer) {
	s.RegisterService(&QuestionerService_ServiceDesc, srv)
}

func _QuestionerService_GetAllQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionerServiceServer).GetAllQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/questioner.QuestionerService/GetAllQuestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionerServiceServer).GetAllQuestions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionerService_CalculateAllQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Answers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionerServiceServer).CalculateAllQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/questioner.QuestionerService/CalculateAllQuestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionerServiceServer).CalculateAllQuestions(ctx, req.(*Answers))
	}
	return interceptor(ctx, in, info, handler)
}

// QuestionerService_ServiceDesc is the grpc.ServiceDesc for QuestionerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuestionerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "questioner.QuestionerService",
	HandlerType: (*QuestionerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllQuestions",
			Handler:    _QuestionerService_GetAllQuestions_Handler,
		},
		{
			MethodName: "CalculateAllQuestions",
			Handler:    _QuestionerService_CalculateAllQuestions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "questioner.proto",
}
