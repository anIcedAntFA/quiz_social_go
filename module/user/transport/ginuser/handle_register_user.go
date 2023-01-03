package ginuser

import (
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"

	hasher "social_quiz/components/hasher"
	userbusiness "social_quiz/module/user/business"
	usermodel "social_quiz/module/user/model"
	userstorage "social_quiz/module/user/storage"

	"github.com/gin-gonic/gin"
)

func HandleRegisterUser(appCtx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		storage := userstorage.NewUserMySQLStorage(db)
		md5 := hasher.NewMd5Hash()
		business := userbusiness.NewRegisterBusiness(storage, md5)

		if err := business.Register(ctx.Request.Context(), &data); err != nil {
			// panic(err)
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		data.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
