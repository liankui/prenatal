// Code generated by chaosmojo. DO NOT EDIT.
// Rerunning chaosmojo will overwrite this file.
// Version: 0.1.0
// Version Date:

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/liankui/prenatal/go/pkg/prenatal"
	"github.com/mojo-lang/core/go/pkg/mojo/core"

	// this service api
	pb "github.com/liankui/prenatal/go/pkg/prenatal/v1"
)

var (
	_ = prenatal.Question{}
	_ = core.Null{}
	_ = prenatal.Answer{}
)

// MakeGRPCServer makes a set of endpoints available as a gRPC QuizServer.
func MakeGRPCServer(endpoints Endpoints, tracer stdopentracing.Tracer, logger log.Logger) pb.QuizServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
		grpctransport.ServerErrorLogger(logger),
	}

	addTracerOption := func(methodName string) []grpctransport.ServerOption {
		if tracer != nil {
			return append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, methodName, logger)))
		}
		return serverOptions
	}

	return &grpcServer{
		// Quiz

		createQuestion: grpctransport.NewServer(
			endpoints.CreateQuestionEndpoint,
			DecodeGRPCCreateQuestionRequest,
			EncodeGRPCCreateQuestionResponse,
			addTracerOption("create_question")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "create_question", logger)))...,
		),
		getQuestion: grpctransport.NewServer(
			endpoints.GetQuestionEndpoint,
			DecodeGRPCGetQuestionRequest,
			EncodeGRPCGetQuestionResponse,
			addTracerOption("get_question")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "get_question", logger)))...,
		),
		updateQuestion: grpctransport.NewServer(
			endpoints.UpdateQuestionEndpoint,
			DecodeGRPCUpdateQuestionRequest,
			EncodeGRPCUpdateQuestionResponse,
			addTracerOption("update_question")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "update_question", logger)))...,
		),
		deleteQuestion: grpctransport.NewServer(
			endpoints.DeleteQuestionEndpoint,
			DecodeGRPCDeleteQuestionRequest,
			EncodeGRPCDeleteQuestionResponse,
			addTracerOption("delete_question")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "delete_question", logger)))...,
		),
		createAnswer: grpctransport.NewServer(
			endpoints.CreateAnswerEndpoint,
			DecodeGRPCCreateAnswerRequest,
			EncodeGRPCCreateAnswerResponse,
			addTracerOption("create_answer")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "create_answer", logger)))...,
		),
	}
}

// grpcServer implements the QuizServer interface
type grpcServer struct {
	pb.UnimplementedQuizServer

	createQuestion grpctransport.Handler
	getQuestion    grpctransport.Handler
	updateQuestion grpctransport.Handler
	deleteQuestion grpctransport.Handler
	createAnswer   grpctransport.Handler
}

// Methods for grpcServer to implement QuizServer interface

func (s *grpcServer) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*prenatal.Question, error) {
	_, rep, err := s.createQuestion.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*prenatal.Question), nil
}

func (s *grpcServer) GetQuestion(ctx context.Context, req *pb.GetQuestionRequest) (*prenatal.Question, error) {
	_, rep, err := s.getQuestion.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*prenatal.Question), nil
}

func (s *grpcServer) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*prenatal.Question, error) {
	_, rep, err := s.updateQuestion.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*prenatal.Question), nil
}

func (s *grpcServer) DeleteQuestion(ctx context.Context, req *pb.DeleteQuestionRequest) (*core.Null, error) {
	_, rep, err := s.deleteQuestion.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*core.Null), nil
}

func (s *grpcServer) CreateAnswer(ctx context.Context, req *pb.CreateAnswerRequest) (*prenatal.Answer, error) {
	_, rep, err := s.createAnswer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*prenatal.Answer), nil
}

// Server Decode

// DecodeGRPCCreateQuestionRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateQuestion request to a user-domain CreateQuestion request. Primarily useful in a server.
func DecodeGRPCCreateQuestionRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateQuestionRequest)
	return req, nil
}

// DecodeGRPCGetQuestionRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetQuestion request to a user-domain GetQuestion request. Primarily useful in a server.
func DecodeGRPCGetQuestionRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetQuestionRequest)
	return req, nil
}

// DecodeGRPCUpdateQuestionRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC UpdateQuestion request to a user-domain UpdateQuestion request. Primarily useful in a server.
func DecodeGRPCUpdateQuestionRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UpdateQuestionRequest)
	return req, nil
}

// DecodeGRPCDeleteQuestionRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC DeleteQuestion request to a user-domain DeleteQuestion request. Primarily useful in a server.
func DecodeGRPCDeleteQuestionRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeleteQuestionRequest)
	return req, nil
}

// DecodeGRPCCreateAnswerRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateAnswer request to a user-domain CreateAnswer request. Primarily useful in a server.
func DecodeGRPCCreateAnswerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateAnswerRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCCreateQuestionResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain CreateQuestion response to a gRPC CreateQuestion reply. Primarily useful in a server.
func EncodeGRPCCreateQuestionResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*prenatal.Question)
	return resp, nil
}

// EncodeGRPCGetQuestionResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain GetQuestion response to a gRPC GetQuestion reply. Primarily useful in a server.
func EncodeGRPCGetQuestionResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*prenatal.Question)
	return resp, nil
}

// EncodeGRPCUpdateQuestionResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain UpdateQuestion response to a gRPC UpdateQuestion reply. Primarily useful in a server.
func EncodeGRPCUpdateQuestionResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*prenatal.Question)
	return resp, nil
}

// EncodeGRPCDeleteQuestionResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain DeleteQuestion response to a gRPC DeleteQuestion reply. Primarily useful in a server.
func EncodeGRPCDeleteQuestionResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*core.Null)
	return resp, nil
}

// EncodeGRPCCreateAnswerResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain CreateAnswer response to a gRPC CreateAnswer reply. Primarily useful in a server.
func EncodeGRPCCreateAnswerResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*prenatal.Answer)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
