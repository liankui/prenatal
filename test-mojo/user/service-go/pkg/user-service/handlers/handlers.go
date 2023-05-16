package handlers

import (
	"context"

	// this service api
	pb "github.com/liankui/prenatal-server/go/pkg/user/v1"
)

type userServer struct {
	pb.UnimplementedUserServer
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.UserServer {
	return userServer{}
}

// GetUserinfo implements Interface.
func (s userServer) GetUserinfo(ctx context.Context, in *pb.GetUserinfoRequest) (*pb.Userinfo, error) {
	resp := &pb.Userinfo{
		// Id:
		// Name:
		// Password:
		// Email:
		// PhoneNumber:
		// Industry:
		// Company:
		// Job:
		// Role:
		// Birthday:
		// Education:
		// Major:
		// Openid:
		// Unionid:
		// Nickname:
		// AvatarUrl:
		// Gender:
		// Country:
		// Province:
		// City:
		// Language:
		// Desc:
		// Source:
		// EventId:
		// InviterId:
		// CreateTime:
		// UpdateTime:
		// DeleteTime:
	}
	return resp, nil
}
