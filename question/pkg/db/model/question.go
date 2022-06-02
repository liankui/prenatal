package model

type Question struct {
	BaseModel
	Desc string `json:"desc" gorm:"column:desc;type:string;size:250;not null;default:'';comment:问题描述"`
	Tips string `json:"tips" gorm:"column:tips;type:string;size:1000;not null;default:'';comment:问题提示"`
}

func (*Question) TableName() string {
	return "question"
}
