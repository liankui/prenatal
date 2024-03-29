// Code generated by mojo. DO NOT EDIT.
// Rerunning mojo will overwrite this file.

package prenatal

import (
	"database/sql/driver"
	"github.com/mojo-lang/db/go/pkg/mojo/db"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func (x QuestionOptions) Value() (driver.Value, error) {
	return db.JSONValuer{}.Value(x)
}

func (x *QuestionOptions) Scan(value interface{}) error {
	return db.JSONScanner{}.Scan(x, value)
}

func (x QuestionOptions) GormDBDataType(gdb *gorm.DB, field *schema.Field) string {
	return db.JSONDbDataType{}.GormDBDataType(gdb, field)
}

func (x QuestionOptions) GormDataType() string {
	return "string"
}
