package questionlikemodel

import (
	"fmt"
	"social_quiz/common"
	"time"
)

const EntityName = "UserLikeQuestion"

type Like struct {
	QuestionId int                `json:"question_id" gorm:"column:question_id;"`
	UserId     int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt  *time.Time         `json:"created_at,omitempty" gorm:"created_at;"`
	User       *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (Like) TableName() string {
	return "question_likes"
}

func (l *Like) GetQuestionId() int {
	return l.QuestionId
}

func ErrorCannotLikeQuestion(err error) *common.AppError {
	return common.NewCustomError(err,
		fmt.Sprintf("cannot like this question"),
		fmt.Sprintf("ErrorCannotLikeQuestion"),
	)
}

func ErrorCannotDislikeQuestion(err error) *common.AppError {
	return common.NewCustomError(err,
		fmt.Sprintf("cannot dislike this question"),
		fmt.Sprintf("ErrorCannotDislikeQuestion"),
	)
}
