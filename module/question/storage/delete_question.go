package questionstorage

import (
	"context"
	"social_quiz/common"
	questionmodel "social_quiz/module/question/model"
)

func (s *questionMySQLStorage) DeleteQuestion(
	ctx context.Context,
	id int,
) error {
	if err := s.db.Table(questionmodel.Question{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
