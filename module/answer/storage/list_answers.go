package answerstorage

import (
	"context"
	"fmt"
	"social_quiz/common"
	answermodel "social_quiz/module/answer/model"
)

func (s *answerMySQLStorage) ListAnswers(
	ctx context.Context,
	param common.PagingParams,
) ([]answermodel.Answer, error) {
	db := s.db.Table(answermodel.Answer{}.TableName())

	if f := param.Search; f != "" {
		//if len(f) > 0 {
		//	db = db.Where("status LIKE ?", fmt.Sprintf("%%%v%%", f))
		//}

		str := fmt.Sprintf("%%%v%%", f)

		if len(f) > 0 {
			db = db.Where("content LIKE ? OR status LIKE ?", str, str)
		}
		//if len(f) > 0 {
		//	db = db.Where("correct in (?)", f)
		//}
	}
	//searchin
	//if err := db.Count(&paging.Total).Error; err != nil {
	//	return nil, common.ErrorDB(err)
	//}
	//
	offset := (param.PageIndex - 1) * param.PageSize

	order := fmt.Sprintf("%v %v", param.SortBy, param.Order)

	var result []answermodel.Answer

	if err := db.
		Offset(offset).
		Limit(param.PageSize).
		Order(order).
		Find(&result).
		Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return result, nil
}
