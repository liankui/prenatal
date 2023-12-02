package model

import (
	"context"
	"sync"

	"github.com/chaos-io/chaos/db"
	"github.com/chaos-io/chaos/logs"

	"github.com/liankui/prenatal/go/pkg/prenatal"
)

var questionModel *QuestionModel
var questionModelOnce sync.Once

type QuestionModel struct {
	DB *db.DB
}

func GetQuestionModel() *QuestionModel {
	questionModelOnce.Do(func() {
		questionModel = NewQuestionModel()
	})

	return questionModel
}

func NewQuestionModel() *QuestionModel {
	m := &QuestionModel{DB: initDB()}
	if !m.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&prenatal.Question{}) {
		if err := d.AutoMigrate(&prenatal.Question{}); err != nil {
			logs.Error("Init Question model err: ", err)
			panic(err)
		}
	}

	return m
}

func (m *QuestionModel) Create(ctx context.Context, question *prenatal.Question) (int64, error) {
	result := m.DB.WithContext(ctx).Create(question)
	return result.RowsAffected, result.Error
}

func (m *QuestionModel) Get(ctx context.Context, id string) (*prenatal.Question, error) {
	question := &prenatal.Question{}
	return question, m.DB.WithContext(ctx).First(question, "id = ?", id).Error
}

func (m *QuestionModel) Delete(ctx context.Context, id string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", id).Delete(&prenatal.Question{})
	return result.RowsAffected, result.Error
}

func (m *QuestionModel) Update(ctx context.Context, question *prenatal.Question) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(question)
	return result.RowsAffected, result.Error
}

func (m *QuestionModel) List(ctx context.Context, filter string, condition ...string) ([]*prenatal.Question, error) {
	var question []*prenatal.Question

	tx := m.DB.WithContext(ctx)
	// todo add condition

	return question, tx.Find(&question).Error
}

func (m *QuestionModel) BatchCreate(ctx context.Context, question ...*prenatal.Question) (int64, error) {
	result := m.DB.WithContext(ctx).CreateInBatches(question, len(question))
	return result.RowsAffected, result.Error
}

func (m *QuestionModel) BatchGet(ctx context.Context, ids ...string) ([]*prenatal.Question, error) {
	var question []*prenatal.Question
	return question, m.DB.WithContext(ctx).Find(&question, "id = ?", ids).Error
}

func (m *QuestionModel) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", ids).Delete(&prenatal.Question{})
	return result.RowsAffected, result.Error
}

func (m *QuestionModel) BatchUpdate(ctx context.Context, question []*prenatal.Question) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(question)
	return result.RowsAffected, result.Error
}
