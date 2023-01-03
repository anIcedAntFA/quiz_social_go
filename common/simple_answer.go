package common

type SimpleAnswer struct {
	SQLModel   `json:",inline"`
	QuestionId int    `json:"-" gorm:"question_id"`
	Content    string `json:"content" gorm:"column:content"`
	Correct    bool   `json:"correct" gorm:"column:correct"`
}

func (ans *SimpleAnswer) TableName() string {
	return "answers"
}

func (ans *SimpleAnswer) Mask(isAdminOrOwner bool) {
	ans.GenerateUID(DbTypeAnswer)
}

type SimpleAnswers []SimpleAnswer
