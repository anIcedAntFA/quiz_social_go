package answermodel

import (
	"fmt"
	"social_quiz/common"
	"strings"
)

const EntityName = "Answer"

type Answer struct {
	common.SQLModel `json:",inline"`
	QuestionId      int    `json:"question_id" gorm:"question_id"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	Content         string `json:"content" gorm:"column:content"`
	Correct         bool   `json:"correct" gorm:"column:correct"`
}

func (Answer) TableName() string {
	return "answers"
}

func (ans *Answer) Mask(isAdminOrOwner bool) {
	ans.GenerateUID(common.DbTypeAnswer)
}

type AnswerCreate struct {
	common.SQLModel `json:",inline"`
	QuestionId      int    `json:"question_id" gorm:"question_id"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	Content         string `json:"content" gorm:"column:content"`
	Correct         bool   `json:"correct" gorm:"column:correct"`
}

func (*AnswerCreate) TableName() string {
	return "answers"
}

func (ans *AnswerCreate) Mask(isAdminOrOwner bool) {
	ans.GenerateUID(common.DbTypeAnswer)
}

func (ans *AnswerCreate) Validate() error {
	dataContent := ans.Content

	if strings.TrimSpace(dataContent) == "" {
		return common.ErrorInvalidRequest(ErrorFieldIsEmpty("answer content"))
	}

	return nil
}

type AnswerUpdate struct {
	common.SQLModel `json:",inline"`
	QuestionId      int    `json:"question_id" gorm:"question_id"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	Content         string `json:"content" gorm:"column:content"`
	Correct         bool   `json:"correct" gorm:"column:correct"`
}

func (*AnswerUpdate) TableName() string {
	return "answers"
}

type Answers []Answer

func ErrorFieldIsEmpty(field string) error {
	return fmt.Errorf("%s cannot be empty", field)
}
