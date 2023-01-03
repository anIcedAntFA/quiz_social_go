package questionstorage

import (
	"context"
	"social_quiz/common"
	questionmodel "social_quiz/module/question/model"
)

func (s *questionMySQLStorage) CreateQuestion(ctx context.Context, data *questionmodel.QuestionCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
