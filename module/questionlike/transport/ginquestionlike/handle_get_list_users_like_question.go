package ginquestionlike

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	questionlikebusiness "social_quiz/module/questionlike/business"
	questionlikemodel "social_quiz/module/questionlike/model"
	questionlikestorage "social_quiz/module/questionlike/storage"
)

func HandleGetListUsersLikeQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		var pagingData common.Paging

		if err := ctx.ShouldBind(&pagingData); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		pagingData.Fulfill()

		filterData := questionlikemodel.Filter{QuestionId: int(uid.GetLocalID())}

		if err := ctx.ShouldBind(&filterData); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := questionlikestorage.NewQuestionLikeSQLStorage(db)
		business := questionlikebusiness.NewListUsersLikeQuestionBusiness(storage)

		result, err := business.ListUsersLikeQuestion(ctx.Request.Context(), &filterData, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filterData))
	}
}
