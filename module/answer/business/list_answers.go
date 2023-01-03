package answerbusiness

import (
	"context"
	"social_quiz/common"
	answermodel "social_quiz/module/answer/model"
)

type ListAnswersStorage interface {
	ListAnswers(
		ctx context.Context,
		filter *answermodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]answermodel.Answer, error)
}

type listAnswersBusiness struct {
	storage ListAnswersStorage
}

func NewListAnswersBusiness(storage ListAnswersStorage) *listAnswersBusiness {
	return &listAnswersBusiness{storage: storage}
}

func (biz *listAnswersBusiness) ListAnswers(
	ctx context.Context,
	filter *answermodel.Filter,
	paging *common.Paging,
) ([]answermodel.Answer, error) {
	result, err := biz.storage.ListAnswers(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrorCannotListEntity(answermodel.EntityName, err)
	}

	return result, nil
}