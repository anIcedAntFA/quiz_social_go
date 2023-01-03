package questionlikestorage

import (
	"context"
	"social_quiz/common"
	questionlikemodel "social_quiz/module/questionlike/model"
)

func (s *questionLikeMySQLStorage) DeleteQuestionLike(ctx context.Context, userId, questionId int) error {
	db := s.db

	if err := db.Table(questionlikemodel.Like{}.TableName()).
		Where("user_id = ? and question_id = ?", userId, questionId).
		Delete(nil).
		Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
