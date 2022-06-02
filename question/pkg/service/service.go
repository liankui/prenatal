package service

import (
	"context"
	"gorm.io/gorm"
)

// QuestionService describes the service.
type QuestionService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicQuestionService struct {
	db *gorm.DB
}

func (b *basicQuestionService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicQuestionService returns a naive, stateless implementation of QuestionService.
func NewBasicQuestionService() QuestionService {
	return &basicQuestionService{}
}

// New returns a QuestionService with all of the expected middleware wired in.
func New(middleware []Middleware) QuestionService {
	var svc QuestionService = NewBasicQuestionService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
