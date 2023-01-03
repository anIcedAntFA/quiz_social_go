package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	userbusiness "social_quiz/module/user/business"
	usermodel "social_quiz/module/user/model"
	userstorage "social_quiz/module/user/storage"
)

func HandleGetListUsers(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := ctx.ShouldBind(&pagingData); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filterData usermodel.Filter

		if err := ctx.ShouldBind(&filterData); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		//filterData.Status = []int{1}

		storage := userstorage.NewUserMySQLStorage(db)
		business := userbusiness.NewListUsersBusiness(storage)

		var users []usermodel.User

		users, err := business.ListQuestions(ctx.Request.Context(), &filterData, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range users {
			users[i].Mask(false)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(users, pagingData, filterData))
	}
}
