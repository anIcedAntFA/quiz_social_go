package questionmodel

import (
	"fmt"
	"social_quiz/common"
	"strings"
)

const EntityName = "Question"

type Question struct {
	common.SQLModel `json:",inline"`
	UserId          int                   `json:"-" gorm:"column:user_id;"`
	Category        string                `json:"category" gorm:"column:category;"`
	Type            string                `json:"type" gorm:"column:type;"`
	Level           string                `json:"level" gorm:"column:level;"`
	Content         string                `json:"content" gorm:"column:content;"`
	Score           int                   `json:"score" gorm:"column:score"`
	Answers         *common.SimpleAnswers `json:"answers" gorm:"preload:false"`
	User            *common.SimpleUser    `json:"user" gorm:"preload:false;"`
	LikeCount       int                   `json:"like_count" gorm:"column:like_count;"`
}

func (Question) TableName() string {
	return "questions"
}

func (q *Question) GetQuestionId() int {
	return q.Id
}

func (q *Question) Mask(isAdminOrOwner bool) {
	q.GenerateUID(common.DbTypeQuestion)

	if u := q.User; u != nil {
		u.Mask(isAdminOrOwner)
	}

	if ans := q.Answers; ans != nil {
		for i := range *ans {
			(*ans)[i].Mask(false)
		}
	}
}

type QuestionCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	Category        string `json:"category" gorm:"column:category;"`
	Type            string `json:"type" gorm:"column:type;"`
	Level           string `json:"level" gorm:"column:level;"`
	Content         string `json:"content" gorm:"column:content;"`
	Score           int    `json:"score" gorm:"column:score"`
}

func (QuestionCreate) TableName() string {
	return Question{}.TableName()
}

func (q *QuestionCreate) Mask(isAdminOrOwner bool) {
	q.GenerateUID(common.DbTypeQuestion)
}

func (q *QuestionCreate) Validate() error {
	dataNames := map[string]string{
		"content":  q.Content,
		"category": q.Category,
		"type":     q.Type,
		"level":    q.Level,
	}

	for k, v := range dataNames {
		v = strings.TrimSpace(v)

		if v == "" {
			return ErrorFieldIsEmpty(k)
		}
	}

	return nil
}

type QuestionUpdate struct {
	common.SQLModel `json:",inline"`
	Category        *string `json:"category" gorm:"column:category;"`
	Type            *string `json:"type" gorm:"column:type;"`
	Level           *string `json:"level" gorm:"column:level;"`
	Content         *string `json:"content" gorm:"column:content;"`
	Score           *int    `json:"score" gorm:"column:score"`
}

func (QuestionUpdate) TableName() string {
	return Question{}.TableName()
}

func ErrorFieldIsEmpty(field string) error {
	return fmt.Errorf("%s cannot be empty", field)
}
