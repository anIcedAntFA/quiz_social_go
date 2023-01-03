package ginuser

import (
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	hasher "social_quiz/components/hasher"
	jwt "social_quiz/components/tokenprovider/jwt"
	userbusiness "social_quiz/module/user/business"
	usermodel "social_quiz/module/user/model"
	userstorage "social_quiz/module/user/storage"

	"github.com/gin-gonic/gin"
)

func HandleLoginUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data usermodel.UserLogin

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		storage := userstorage.NewUserMySQLStorage(db)
		md5 := hasher.NewMd5Hash()
		business := userbusiness.NewLoginBusiness(storage, tokenProvider, md5, 60*60*24*30)

		account, err := business.Login(ctx.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
