package answerstorage

import (
	"context"
	"social_quiz/common"
	answermodel "social_quiz/module/answer/model"
)

func (s *answerMySQLStorage) UpdateAnswer(
	ctx context.Context,
	id int,
	data *answermodel.AnswerUpdate,
) error {
	if err := s.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
