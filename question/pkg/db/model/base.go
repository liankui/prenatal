package model

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	Id        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
