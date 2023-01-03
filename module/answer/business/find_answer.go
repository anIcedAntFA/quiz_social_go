package answerbusiness

import (
	"context"
	"gorm.io/gorm"
	"social_quiz/common"
	answermodel "social_quiz/module/answer/model"
)

type FindAnswerStorage interface {
	FindAnswer(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*answermodel.Answer, error)
}

type findAnswerBusiness struct {
	storage FindAnswerStorage
}

func NewFindAnswerBusiness(storage FindAnswerStorage) *findAnswerBusiness {
	return &findAnswerBusiness{storage: storage}
}

func (biz *findAnswerBusiness) FindAnswer(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*answermodel.Answer, error) {
	result, err := biz.storage.FindAnswer(ctx, condition)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrorRecordNotFound
		}

		return nil, common.ErrorEntityNotFound(answermodel.EntityName, err)
	}

	return result, err
}
