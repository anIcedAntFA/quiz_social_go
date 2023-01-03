package questionstorage

import (
	"context"
	"gorm.io/gorm"
	"social_quiz/common"
	questionmodel "social_quiz/module/question/model"
)

func (s *questionMySQLStorage) FindQuestion(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*questionmodel.Question, error) {
	db := s.db

	for _, v := range moreKeys {
		db = db.Preload(v)
	}

	var result questionmodel.Question

	if err := db.Where(condition).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrorRecordNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &result, nil
}
