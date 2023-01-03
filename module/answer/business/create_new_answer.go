package answerbusiness

import (
	"context"
	"social_quiz/common"
	answermodel "social_quiz/module/answer/model"
)

type CreateAnswerStorage interface {
	CreateAnswer(ctx context.Context, data *answermodel.AnswerCreate) error
}

type createAnswerBusiness struct {
	storage CreateAnswerStorage
}

func NewCreateAnswerBusiness(storage CreateAnswerStorage) *createAnswerBusiness {
	return &createAnswerBusiness{storage: storage}
}

func (biz *createAnswerBusiness) CreateAnswer(ctx context.Context, data *answermodel.AnswerCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrorInvalidRequest(err)
	}

	if err := biz.storage.CreateAnswer(ctx, data); err != nil {
		return common.ErrorCannotCreateEntity(answermodel.EntityName, err)
	}

	return nil
}
