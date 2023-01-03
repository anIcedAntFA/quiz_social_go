package answerbusiness

import (
	"context"
	"social_quiz/common"
	answermodel "social_quiz/module/answer/model"
)

type UpdateAnswerStorage interface {
	FindAnswer(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*answermodel.Answer, error)

	UpdateAnswer(ctx context.Context, id int, data *answermodel.AnswerUpdate) error
}

type updateAnswerBusiness struct {
	storage UpdateAnswerStorage
}

func NewUpdateAnswerBusiness(storage UpdateAnswerStorage) *updateAnswerBusiness {
	return &updateAnswerBusiness{storage: storage}
}

func (biz *updateAnswerBusiness) UpdateAnswer(
	ctx context.Context,
	id int,
	data *answermodel.AnswerUpdate,
) error {
	oldData, err := biz.storage.FindAnswer(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrorEntityNotFound(answermodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrorEntityDeleted(answermodel.EntityName, err)
	}

	if err := biz.storage.UpdateAnswer(ctx, id, data); err != nil {
		return common.ErrorCannotUpdateEntity(answermodel.EntityName, err)
	}

	return nil
}
