// Code generated by ncraft. DO NOT EDIT.
// Rerunning ncraft will overwrite this file.
// Version: 0.0.1
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

	// this service api
	pb "git.company.com/examples/hello-world/go/pkg/hello-world/v1"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC HelloWorldServer.
func MakeGRPCServer(endpoints Endpoints, tracer stdopentracing.Tracer, logger log.Logger) pb.HelloWorldServer {
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
		// HelloWorld

		getEcho: grpctransport.NewServer(
			endpoints.GetEchoEndpoint,
			DecodeGRPCGetEchoRequest,
			EncodeGRPCGetEchoResponse,
			addTracerOption("get_echo")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "get_echo", logger)))...,
		),
	}
}

// grpcServer implements the HelloWorldServer interface
type grpcServer struct {
	pb.UnimplementedHelloWorldServer

	getEcho grpctransport.Handler
}

// Methods for grpcServer to implement HelloWorldServer interface

func (s *grpcServer) GetEcho(ctx context.Context, req *pb.GetEchoRequest) (*pb.Echo, error) {
	_, rep, err := s.getEcho.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Echo), nil
}

// Server Decode

// DecodeGRPCGetEchoRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetEcho request to a user-domain GetEcho request. Primarily useful in a server.
func DecodeGRPCGetEchoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetEchoRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCGetEchoResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain GetEcho response to a gRPC GetEcho reply. Primarily useful in a server.
func EncodeGRPCGetEchoResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.Echo)
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