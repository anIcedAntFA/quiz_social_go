package userbusiness

import (
	"context"
	"social_quiz/common"
	"social_quiz/components/tokenprovider"
	usermodel "social_quiz/module/user/model"
)

type LoginStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
}

type loginBusiness struct {
	storage       LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(
	storage LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	expiry int,
) *loginBusiness {
	return &loginBusiness{
		storage:       storage,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. find user, email
// 2. hash password from input and compare with password in db
// 3. provider: issue JWT token for client
// 		access token and refresh token
// 4. return token(s)

func (biz *loginBusiness) Login(
	ctx context.Context,
	data *usermodel.UserLogin,
) (*tokenprovider.Token, error) {
	user, err := biz.storage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrorEmailOrPasswordInvalid
	}

	hashedPassword := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != hashedPassword {
		return nil, usermodel.ErrorEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrorInternal(err)
	}

	return accessToken, nil
}
