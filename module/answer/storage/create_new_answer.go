package answerstorage

import (
	"context"
	"social_quiz/common"
	answermodel "social_quiz/module/answer/model"
)

func (s *answerMySQLStorage) CreateAnswer(ctx context.Context, data *answermodel.AnswerCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
