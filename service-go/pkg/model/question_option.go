package model

import (
	"context"
	"sync"

	"github.com/chaos-io/chaos/db"
	"github.com/chaos-io/chaos/logs"

	"github.com/liankui/prenatal/go/pkg/prenatal"
)

var questionOptionModel *QuestionOptionModel
var questionOptionModelOnce sync.Once

type QuestionOptionModel struct {
	DB *db.DB
}

func GetQuestionOptionModel() *QuestionOptionModel {
	questionOptionModelOnce.Do(func() {
		questionOptionModel = NewQuestionOptionModel()
	})

	return questionOptionModel
}

func NewQuestionOptionModel() *QuestionOptionModel {
	m := &QuestionOptionModel{DB: initDB()}
	if !m.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&prenatal.QuestionOption{}) {
		if err := d.AutoMigrate(&prenatal.QuestionOption{}); err != nil {
			logs.Error("Init QuestionOption model error: ", err)
			panic(err)
		}
	}

	return m
}

func (m *QuestionOptionModel) Create(ctx context.Context, questionOption *prenatal.QuestionOption) (int64, error) {
	result := m.DB.WithContext(ctx).Create(questionOption)
	return result.RowsAffected, result.Error
}

func (m *QuestionOptionModel) Get(ctx context.Context, id string) (*prenatal.QuestionOption, error) {
	questionOption := &prenatal.QuestionOption{}
	return questionOption, m.DB.WithContext(ctx).First(questionOption, "id = ?", id).Error
}

func (m *QuestionOptionModel) Delete(ctx context.Context, id string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", id).Delete(&prenatal.QuestionOption{})
	return result.RowsAffected, result.Error
}

func (m *QuestionOptionModel) Update(ctx context.Context, questionOption *prenatal.QuestionOption) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(questionOption)
	return result.RowsAffected, result.Error
}

func (m *QuestionOptionModel) List(ctx context.Context, filter string, condition ...string) ([]*prenatal.QuestionOption, error) {
	var questionOption []*prenatal.QuestionOption

	tx := m.DB.WithContext(ctx)
	// todo add condition

	return questionOption, tx.Find(&questionOption).Error
}

func (m *QuestionOptionModel) BatchCreate(ctx context.Context, questionOption ...*prenatal.QuestionOption) (int64, error) {
	result := m.DB.WithContext(ctx).CreateInBatches(questionOption, len(questionOption))
	return result.RowsAffected, result.Error
}

func (m *QuestionOptionModel) BatchGet(ctx context.Context, ids ...string) ([]*prenatal.QuestionOption, error) {
	var questionOption []*prenatal.QuestionOption
	return questionOption, m.DB.WithContext(ctx).Find(&questionOption, "id = ?", ids).Error
}

func (m *QuestionOptionModel) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", ids).Delete(&prenatal.QuestionOption{})
	return result.RowsAffected, result.Error
}

func (m *QuestionOptionModel) BatchUpdate(ctx context.Context, questionOption []*prenatal.QuestionOption) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(questionOption)
	return result.RowsAffected, result.Error
}
