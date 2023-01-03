package questionbusiness

import (
	"context"
	"errors"
	"social_quiz/common"
	questionmodel "social_quiz/module/question/model"
)

type UpdateQuestionStorage interface {
	FindQuestion(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*questionmodel.Question, error)

	UpdateQuestion(ctx context.Context, id int, data *questionmodel.QuestionUpdate) error
}

type updateQuestionBusiness struct {
	storage   UpdateQuestionStorage
	requester common.Requester
}

func NewUpdateQuestionBusiness(storage UpdateQuestionStorage, requester common.Requester) *updateQuestionBusiness {
	return &updateQuestionBusiness{
		storage:   storage,
		requester: requester,
	}
}

func (biz *updateQuestionBusiness) UpdateQuestion(
	ctx context.Context,
	id int,
	data *questionmodel.QuestionUpdate,
) error {
	oldData, err := biz.storage.FindQuestion(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrorEntityNotFound(questionmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrorEntityDeleted(questionmodel.EntityName, err)
	}

	if biz.requester.GetUserId() != oldData.UserId || biz.requester.GetRole() != "admin" {
		return common.ErrorNoPermission(errors.New("you have no permission to update this question"))
	}

	if err := biz.storage.UpdateQuestion(ctx, id, data); err != nil {
		return common.ErrorCannotUpdateEntity(questionmodel.EntityName, err)
	}

	return nil
}
