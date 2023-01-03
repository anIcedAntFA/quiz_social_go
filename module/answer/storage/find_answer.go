package answerstorage

import (
	"context"
	"gorm.io/gorm"
	"social_quiz/common"
	answermodel "social_quiz/module/answer/model"
)

func (s *answerMySQLStorage) FindAnswer(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*answermodel.Answer, error) {
	var result answermodel.Answer

	if err := s.db.Where(condition).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrorRecordNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &result, nil
}
