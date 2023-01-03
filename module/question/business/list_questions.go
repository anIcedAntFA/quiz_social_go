package questionbusiness

import (
	"context"
	"social_quiz/common"
	questionmodel "social_quiz/module/question/model"
)

type ListQuestionsRepository interface {
	ListQuestions(
		ctx context.Context,
		filter *questionmodel.Filter,
		paging *common.Paging,
	) ([]questionmodel.Question, error)
}

type listQuestionsBusiness struct {
	repo ListQuestionsRepository
}

func NewListQuestionsBusiness(repo ListQuestionsRepository) *listQuestionsBusiness {
	return &listQuestionsBusiness{repo: repo}
}

func (biz listQuestionsBusiness) ListQuestions(
	ctx context.Context,
	filter *questionmodel.Filter,
	paging *common.Paging,
) ([]questionmodel.Question, error) {
	result, err := biz.repo.ListQuestions(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrorCannotListEntity(questionmodel.EntityName, err)
	}

	return result, nil
}
