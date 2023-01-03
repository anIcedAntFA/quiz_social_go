package userbusiness

import (
	"context"
	"social_quiz/common"
	usermodel "social_quiz/module/user/model"
)

type RegisterStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)

	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	storage RegisterStorage
	hasher  Hasher
}

func NewRegisterBusiness(storage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		storage: storage,
		hasher:  hasher,
	}
}

func (biz *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.storage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		if user.Status == 0 {
			return usermodel.ErrorUserDisabledOrBanned
		}

		return usermodel.ErrorEmailExisted
	}

	salt := common.GenerateSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" //hard code

	if err := biz.storage.CreateUser(ctx, data); err != nil {
		return common.ErrorCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
