package answerbusiness

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"social_quiz/common"
	answermodel "social_quiz/module/answer/model"
)

type DeleteAnswerStorage interface {
	FindAnswer(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*answermodel.Answer, error)

	DeleteAnswer(ctx context.Context, id int) error
}

type deleteAnswerBusiness struct {
	storage   DeleteAnswerStorage
	requester common.Requester
}

func NewDeleteAnswerBusiness(storage DeleteAnswerStorage, requester common.Requester) *deleteAnswerBusiness {
	return &deleteAnswerBusiness{
		storage:   storage,
		requester: requester,
	}
}

func (biz *deleteAnswerBusiness) DeleteAnswer(ctx context.Context, id int) error {
	oldData, err := biz.storage.FindAnswer(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrorRecordNotFound
		}

		return common.ErrorEntityNotFound(answermodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrorEntityDeleted(answermodel.EntityName, nil)
	}

	if biz.requester.GetUserId() != oldData.UserId || biz.requester.GetRole() != "admin" {
		return common.ErrorNoPermission(errors.New("you have no permission to delete this answer"))
	}

	if err := biz.storage.DeleteAnswer(ctx, id); err != nil {
		return common.ErrorCannotDeleteEntity(answermodel.EntityName, err)
	}

	return nil
}
