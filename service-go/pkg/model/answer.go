package model

import (
	"context"
	"sync"

	"github.com/chaos-io/chaos/db"
	"github.com/chaos-io/chaos/logs"

	"github.com/liankui/prenatal/go/pkg/prenatal"
)

var answerModel *AnswerModel
var answerModelOnce sync.Once

type AnswerModel struct {
	DB *db.DB
}

func GetAnswerModel() *AnswerModel {
	answerModelOnce.Do(func() {
		answerModel = NewAnswerModel()
	})

	return answerModel
}

func NewAnswerModel() *AnswerModel {
	m := &AnswerModel{DB: initDB()}
	if !m.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&prenatal.Answer{}) {
		if err := d.AutoMigrate(&prenatal.Answer{}); err != nil {
			logs.Error("Init Answer model err: ", err)
			panic(err)
		}
	}

	return m
}

func (m *AnswerModel) Create(ctx context.Context, answer *prenatal.Answer) (int64, error) {
	result := m.DB.WithContext(ctx).Create(answer)
	return result.RowsAffected, result.Error
}

func (m *AnswerModel) Get(ctx context.Context, id string) (*prenatal.Answer, error) {
	answer := &prenatal.Answer{}
	return answer, m.DB.WithContext(ctx).First(answer, "id = ?", id).Error
}

func (m *AnswerModel) Delete(ctx context.Context, id string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", id).Delete(&prenatal.Answer{})
	return result.RowsAffected, result.Error
}

func (m *AnswerModel) Update(ctx context.Context, answer *prenatal.Answer) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(answer)
	return result.RowsAffected, result.Error
}

func (m *AnswerModel) List(ctx context.Context, filter string, condition ...string) ([]*prenatal.Answer, error) {
	var answer []*prenatal.Answer

	tx := m.DB.WithContext(ctx)
	// todo add condition

	return answer, tx.Find(&answer).Error
}

func (m *AnswerModel) BatchCreate(ctx context.Context, answer ...*prenatal.Answer) (int64, error) {
	result := m.DB.WithContext(ctx).CreateInBatches(answer, len(answer))
	return result.RowsAffected, result.Error
}

func (m *AnswerModel) BatchGet(ctx context.Context, ids ...string) ([]*prenatal.Answer, error) {
	var answer []*prenatal.Answer
	return answer, m.DB.WithContext(ctx).Find(&answer, "id = ?", ids).Error
}

func (m *AnswerModel) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", ids).Delete(&prenatal.Answer{})
	return result.RowsAffected, result.Error
}

func (m *AnswerModel) BatchUpdate(ctx context.Context, answer []*prenatal.Answer) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(answer)
	return result.RowsAffected, result.Error
}
