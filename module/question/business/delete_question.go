package questionbusiness

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"social_quiz/common"
	questionmodel "social_quiz/module/question/model"
)

type DeleteQuestionStorage interface {
	FindQuestion(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*questionmodel.Question, error)

	DeleteQuestion(ctx context.Context, id int) error
}

type deleteQuestionBusiness struct {
	storage   DeleteQuestionStorage
	requester common.Requester
}

func NewDeleteQuestionBusiness(storage DeleteQuestionStorage, requester common.Requester) *deleteQuestionBusiness {
	return &deleteQuestionBusiness{
		storage:   storage,
		requester: requester,
	}
}

func (biz *deleteQuestionBusiness) DeleteQuestion(ctx context.Context, id int) error {
	oldData, err := biz.storage.FindQuestion(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrorRecordNotFound
		}

		return common.ErrorEntityNotFound(questionmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrorEntityDeleted(questionmodel.EntityName, nil)
	}

	if biz.requester.GetUserId() != oldData.UserId || biz.requester.GetRole() != "admin" {
		return common.ErrorNoPermission(errors.New("you have no permission to delete this question"))
	}

	if err := biz.storage.DeleteQuestion(ctx, id); err != nil {
		return common.ErrorCannotDeleteEntity(questionmodel.EntityName, err)
	}

	return nil
}
