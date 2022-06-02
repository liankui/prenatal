package model

type QuestionOption struct {
	BaseModel
	QuestionId int    `json:"question_id" gorm:"column:question_id;type:int;not null;default:0"`
	IsAnswer   bool   `json:"is_answer" gorm:"column:is_answer;type:bool;comment:是否为正确答案"`
	Desc       string `json:"desc" gorm:"column:desc;type:string;size:250;not null;default:'';comment:该选项的描述"`
}

func (*QuestionOption) TableName() string {
	return "question_option"
}
