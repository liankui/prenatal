package model

/*
import (
	"context"
	"sync"

	"github.com/chaos-io/chaos/db"
	"github.com/chaos-io/chaos/logs"

    // TODO:
    // 1. rename this filename to {entity}.go
    // 2. cancel the code comment, update the following import and use "github.com/{repo}/{package}/go/pkg/{package}"
    // 3. replace data to {package}, replace entity to {entity}, replace Entity to {Entity}
	"github.com/{repo}/data/go/pkg/data"
)

var entityModel *EntityModel
var entityModelOnce sync.Once

type EntityModel struct {
	DB *db.DB
}

func GetEntityModel() *EntityModel {
	entityModelOnce.Do(func() {
		entityModel = NewEntityModel()
	})

	return entityModel
}

func NewEntityModel() *EntityModel {
	m := &EntityModel{DB: initDB()}
	if !m.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&data.Entity{}) {
		if err := d.AutoMigrate(&data.Entity{}); err != nil {
			logs.Error("Init Entity model err: ", err)
			panic(err)
		}
	}

	return m
}

func (m *EntityModel) Create(ctx context.Context, entity *data.Entity) (int64, error) {
	result := m.DB.WithContext(ctx).Create(entity)
	return result.RowsAffected, result.Error
}

func (m *EntityModel) Get(ctx context.Context, id string) (*data.Entity, error) {
	entity := &data.Entity{}
	return entity, m.DB.WithContext(ctx).First(entity, "id = ?", id).Error
}

func (m *EntityModel) Delete(ctx context.Context, id string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", id).Delete(&data.Entity{})
	return result.RowsAffected, result.Error
}

func (m *EntityModel) Update(ctx context.Context, entity *data.Entity) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(entity)
	return result.RowsAffected, result.Error
}

func (m *EntityModel) List(ctx context.Context, filter string, condition ...string) ([]*data.Entity, error) {
	var entity []*data.Entity

	tx := m.DB.WithContext(ctx)
	// todo add condition

	return entity, tx.Find(&entity).Error
}

func (m *EntityModel) BatchCreate(ctx context.Context, entity ...*data.Entity) (int64, error) {
	result := m.DB.WithContext(ctx).CreateInBatches(entity, len(entity))
	return result.RowsAffected, result.Error
}

func (m *EntityModel) BatchGet(ctx context.Context, ids ...string) ([]*data.Entity, error) {
	var entity []*data.Entity
	return entity, m.DB.WithContext(ctx).Find(&entity, "id = ?", ids).Error
}

func (m *EntityModel) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", ids).Delete(&data.Entity{})
	return result.RowsAffected, result.Error
}

func (m *EntityModel) BatchUpdate(ctx context.Context, entity []*data.Entity) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(entity)
	return result.RowsAffected, result.Error
}
*/
