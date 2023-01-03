package questionbusiness

import (
	"context"
	"social_quiz/common"
	questionmodel "social_quiz/module/question/model"
)

type CreateQuestionStorage interface {
	CreateQuestion(ctx context.Context, data *questionmodel.QuestionCreate) error
}

type createQuestionBusiness struct {
	storage CreateQuestionStorage
}

func NewCreateQuestionBusiness(storage CreateQuestionStorage) *createQuestionBusiness {
	return &createQuestionBusiness{storage: storage}
}

func (biz *createQuestionBusiness) CreateQuestion(ctx context.Context, data *questionmodel.QuestionCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrorInvalidRequest(err)
	}

	if err := biz.storage.CreateQuestion(ctx, data); err != nil {
		return common.ErrorCannotCreateEntity(questionmodel.EntityName, err)
	}

	return nil
}
