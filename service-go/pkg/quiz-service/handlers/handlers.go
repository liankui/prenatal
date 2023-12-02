package handlers

import (
	"context"

	"github.com/chaos-io/chaos/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/segmentio/ksuid"

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
	logs.Infow("CreateQuestion", "request", in)

	if in.Question.Question == "" {
		return nil, core.NewErrorFrom(400, "question is NOT set")
	}
	if len(in.Question.Options) == 0 {
		return nil, core.NewErrorFrom(400, "options is NOT set")
	}

	if in.Question.Id == "" {
		in.Question.Id = ksuid.New().String()
	}

	in.Question.CreateTime = core.Now()

	_, err := model.GetQuestionModel().Create(ctx, in.Question)
	if err != nil {
		return nil, core.NewErrorFrom(500, "failed to insert to question model, error: %v", err)
	}

	// add question options
	options := in.Question.Options
	for _, op := range options {
		op.Id = ksuid.New().String()
		op.QuestionId = in.Question.Id
		op.CreateTime = core.Now()
	}

	_, err = model.GetQuestionOptionModel().BatchCreate(ctx, options...)
	if err != nil {
		return nil, core.NewErrorFrom(500, "failed to insert to question option model, error: %v", err)
	}

	return in.Question, nil
}

// GetQuestion implements Interface.
func (s quizServer) GetQuestion(ctx context.Context, in *pb.GetQuestionRequest) (*prenatal.Question, error) {
	logs.Infow("GetQuestion", "request", in)

	if in.Id == "" {
		return nil, core.NewErrorFrom(400, "id is NOT set")
	}

	que, err := model.GetQuestionModel().Get(ctx, in.Id)
	if err != nil {
		return nil, core.NewErrorFrom(500, "failed to get question, error: %v", err)
	}

	// avoid to expose the right answer
	for _, op := range que.GetOptions() {
		op.IsCorrect = false
	}

	return que, nil
}

// UpdateQuestion implements Interface.
func (s quizServer) UpdateQuestion(ctx context.Context, in *pb.UpdateQuestionRequest) (*prenatal.Question, error) {
	logs.Infow("UpdateQuestion", "request", in)

	if in.Question.Id == "" {
		return nil, core.NewErrorFrom(400, "id is NOT set")
	}
	if in.Question.Question == "" {
		return nil, core.NewErrorFrom(400, "question is NOT set")
	}

	_, err := model.GetQuestionModel().Update(ctx, in.Question)
	if err != nil {
		return nil, core.NewErrorFrom(500, "failed to update question, error: %v", err)
	}

	return in.Question, nil
}

// DeleteQuestion implements Interface.
func (s quizServer) DeleteQuestion(ctx context.Context, in *pb.DeleteQuestionRequest) (*core.Null, error) {
	logs.Infow("DeleteQuestion", "request", in)

	if in.Id == "" {
		return nil, core.NewErrorFrom(400, "id is NOT set")
	}

	_, err := model.GetQuestionModel().Delete(ctx, in.Id)
	if err != nil {
		return nil, core.NewErrorFrom(500, "failed to delete question, error: %v", err)
	}

	resp := &core.Null{}
	return resp, nil
}

// CreateAnswer implements Interface.
func (s quizServer) CreateAnswer(ctx context.Context, in *pb.CreateAnswerRequest) (*prenatal.Answer, error) {
	logs.Infow("CreateAnswer", "request", in)

	resp := &prenatal.Answer{
		// Id:
		// QuestionId:
		// QuestionCategory:
		// Answer:
		// CorrectAnswer:
		// UserId:
		// UserName:
		// CreateTime:
		// UpdateTime:
	}
	return resp, nil
}
