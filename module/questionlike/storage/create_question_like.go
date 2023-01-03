package questionlikestorage

import (
	"context"
	"social_quiz/common"
	questionlikemodel "social_quiz/module/questionlike/model"
)

func (s *questionLikeMySQLStorage) CreateQuestionLike(ctx context.Context, data *questionlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
