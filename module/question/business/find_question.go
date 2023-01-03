package questionbusiness

import (
	"context"
	"gorm.io/gorm"
	"social_quiz/common"
	questionmodel "social_quiz/module/question/model"
)

type FindQuestionStorage interface {
	FindQuestion(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*questionmodel.Question, error)
}

type findQuestionBusiness struct {
	storage FindQuestionStorage
}

func NewFindQuestionBusiness(storage FindQuestionStorage) *findQuestionBusiness {
	return &findQuestionBusiness{storage: storage}
}

func (biz findQuestionBusiness) FindQuestion(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*questionmodel.Question, error) {
	result, err := biz.storage.FindQuestion(ctx, condition, "Answers", "User")

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrorRecordNotFound
		}

		return nil, common.ErrorEntityNotFound(questionmodel.EntityName, err)
	}

	return result, err
}
