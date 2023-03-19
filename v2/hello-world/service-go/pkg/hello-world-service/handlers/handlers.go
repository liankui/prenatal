package handlers

import (
	"context"

	// this service api
	pb "git.company.com/examples/hello-world/go/pkg/hello-world/v1"
)

type helloWorldServer struct {
	pb.UnimplementedHelloWorldServer
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.HelloWorldServer {
	return helloWorldServer{}
}

func (s helloWorld) GetEcho(ctx context.Context, in *pb.GetEchoRequest) (*pb.Echo, error) {
	resp := pb.Echo{
		Name:    in.Name,
		Message: "Hello, " + in.Name + "!",
	}
	return &resp, nil
}
