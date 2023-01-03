package questionstorage

import (
	"context"
	"gorm.io/gorm"
	"social_quiz/common"
	questionmodel "social_quiz/module/question/model"
)

func (s *questionMySQLStorage) UpdateQuestion(
	ctx context.Context,
	id int,
	data *questionmodel.QuestionUpdate,
) error {
	if err := s.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}

func (s *questionMySQLStorage) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(questionmodel.Question{}.TableName()).
		Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count + ?", 1)).
		Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}

func (s *questionMySQLStorage) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(questionmodel.Question{}.TableName()).
		Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count - ?", 1)).
		Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
