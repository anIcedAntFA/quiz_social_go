package userbusiness

import (
	"context"
	"social_quiz/common"
	usermodel "social_quiz/module/user/model"
)

type ListUsersStorage interface {
	ListUsers(
		ctx context.Context,
		filter *usermodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]usermodel.User, error)
}

type listUsersBusiness struct {
	storage ListUsersStorage
}

func NewListUsersBusiness(storage ListUsersStorage) *listUsersBusiness {
	return &listUsersBusiness{
		storage: storage,
	}
}

func (biz listUsersBusiness) ListQuestions(
	ctx context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
) ([]usermodel.User, error) {
	result, err := biz.storage.ListUsers(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrorCannotListEntity(usermodel.EntityName, err)
	}

	return result, nil
}
