package model

import (
	"context"
	"sync"

	"github.com/chaos-io/chaos/db"
	"github.com/chaos-io/chaos/logs"

	"github.com/liankui/prenatal/go/pkg/prenatal"
)

var commentModel *CommentModel
var commentModelOnce sync.Once

type CommentModel struct {
	DB *db.DB
}

func GetCommentModel() *CommentModel {
	commentModelOnce.Do(func() {
		commentModel = NewCommentModel()
	})

	return commentModel
}

func NewCommentModel() *CommentModel {
	m := &CommentModel{DB: initDB()}
	if !m.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&prenatal.Comment{}) {
		if err := d.AutoMigrate(&prenatal.Comment{}); err != nil {
			logs.Error("Init Comment model err: ", err)
			panic(err)
		}
	}

	return m
}

func (m *CommentModel) Create(ctx context.Context, comment *prenatal.Comment) (int64, error) {
	result := m.DB.WithContext(ctx).Create(comment)
	return result.RowsAffected, result.Error
}

func (m *CommentModel) Get(ctx context.Context, id string) (*prenatal.Comment, error) {
	comment := &prenatal.Comment{}
	return comment, m.DB.WithContext(ctx).First(comment, "id = ?", id).Error
}

func (m *CommentModel) Delete(ctx context.Context, id string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", id).Delete(&prenatal.Comment{})
	return result.RowsAffected, result.Error
}

func (m *CommentModel) Update(ctx context.Context, comment *prenatal.Comment) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(comment)
	return result.RowsAffected, result.Error
}

func (m *CommentModel) List(ctx context.Context, filter string, condition ...string) ([]*prenatal.Comment, error) {
	var comment []*prenatal.Comment

	tx := m.DB.WithContext(ctx)
	// todo add condition

	return comment, tx.Find(&comment).Error
}

func (m *CommentModel) BatchCreate(ctx context.Context, comment ...*prenatal.Comment) (int64, error) {
	result := m.DB.WithContext(ctx).CreateInBatches(comment, len(comment))
	return result.RowsAffected, result.Error
}

func (m *CommentModel) BatchGet(ctx context.Context, ids ...string) ([]*prenatal.Comment, error) {
	var comment []*prenatal.Comment
	return comment, m.DB.WithContext(ctx).Find(&comment, "id = ?", ids).Error
}

func (m *CommentModel) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", ids).Delete(&prenatal.Comment{})
	return result.RowsAffected, result.Error
}

func (m *CommentModel) BatchUpdate(ctx context.Context, comment []*prenatal.Comment) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(comment)
	return result.RowsAffected, result.Error
}
