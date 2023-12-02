package handlers

import (
	"context"

	"github.com/mojo-lang/core/go/pkg/mojo/core"

	"github.com/liankui/prenatal/go/pkg/prenatal"
	"github.com/liankui/prenatal/service-go/pkg/model"

	// this service api
	pb "github.com/liankui/prenatal/go/pkg/prenatal/v1"
)

var (
	_ = prenatal.Question{}
	_ = core.Null{}
)

type quizServer struct {
	pb.UnimplementedQuizServer
}

func init() {
	model.InitModel()
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.QuizServer {
	return quizServer{}
}

// CreateQuestion implements Interface.
func (s quizServer) CreateQuestion(ctx context.Context, in *pb.CreateQuestionRequest) (*prenatal.Question, error) {
	resp := &prenatal.Question{
		// Id:
		// Category:
		// Question:
		// Options:
		// CorrectAnswer:
		// Explanation:
		// CreateTime:
		// UpdateTime:
	}
	return resp, nil
}

// GetQuestion implements Interface.
func (s quizServer) GetQuestion(ctx context.Context, in *pb.GetQuestionRequest) (*prenatal.Question, error) {
	resp := &prenatal.Question{
		// Id:
		// Category:
		// Question:
		// Options:
		// CorrectAnswer:
		// Explanation:
		// CreateTime:
		// UpdateTime:
	}
	return resp, nil
}

// UpdateQuestion implements Interface.
func (s quizServer) UpdateQuestion(ctx context.Context, in *pb.UpdateQuestionRequest) (*prenatal.Question, error) {
	resp := &prenatal.Question{
		// Id:
		// Category:
		// Question:
		// Options:
		// CorrectAnswer:
		// Explanation:
		// CreateTime:
		// UpdateTime:
	}
	return resp, nil
}

// DeleteQuestion implements Interface.
func (s quizServer) DeleteQuestion(ctx context.Context, in *pb.DeleteQuestionRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}
