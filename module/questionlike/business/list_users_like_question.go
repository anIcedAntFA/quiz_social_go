package questionlikebusiness

import (
	"context"
	"social_quiz/common"
	questionlikemodel "social_quiz/module/questionlike/model"
)

type ListUsersLikeQuestionStorage interface {
	ListUsersLikeQuestion(
		ctx context.Context,
		conditions map[string]interface{},
		filter *questionlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUsersLikeQuestionBusiness struct {
	storage ListUsersLikeQuestionStorage
}

func NewListUsersLikeQuestionBusiness(storage ListUsersLikeQuestionStorage) *listUsersLikeQuestionBusiness {
	return &listUsersLikeQuestionBusiness{storage: storage}
}

func (biz *listUsersLikeQuestionBusiness) ListUsersLikeQuestion(
	ctx context.Context,
	filter *questionlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	result, err := biz.storage.ListUsersLikeQuestion(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrorCannotListEntity(questionlikemodel.EntityName, err)
	}

	return result, nil
}
