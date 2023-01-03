package answerstorage

import (
	"context"
	"social_quiz/common"
	answermodel "social_quiz/module/answer/model"
)

func (s *answerMySQLStorage) DeleteAnswer(ctx context.Context, id int) error {
	if err := s.db.Table(answermodel.Answer{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
