package answerstorage

import (
	"context"
	"social_quiz/common"
	answermodel "social_quiz/module/answer/model"
)

func (s *answerMySQLStorage) ListAnswers(
	ctx context.Context,
	filter *answermodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]answermodel.Answer, error) {
	db := s.db.Table(answermodel.Answer{}.TableName())

	if f := filter; f != nil {
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
		if len(f.Content) > 0 {
			db = db.Where("status in (?)", f.Content)
		}
		if len(f.Correct) > 0 {
			db = db.Where("status in (?)", f.Correct)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit

	var result []answermodel.Answer

	if err := db.
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).
		Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return result, nil
}
