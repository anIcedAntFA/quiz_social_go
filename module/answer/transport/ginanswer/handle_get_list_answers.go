package ginanswer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	answerbusiness "social_quiz/module/answer/business"
	answermodel "social_quiz/module/answer/model"
	answerstorage "social_quiz/module/answer/storage"
)

func HandleGetListAnswers(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pagingData common.Paging

		if err := ctx.ShouldBind(&pagingData); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filterData answermodel.Filter

		if err := ctx.ShouldBind(&filterData); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		filterData.Status = []int{1}

		db := appCtx.GetMainDBConnection()

		storage := answerstorage.NewAnswerMySQLStorage(db)
		business := answerbusiness.NewListAnswersBusiness(storage)

		answers, err := business.ListAnswers(ctx.Request.Context(), &filterData, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range answers {
			answers[i].Mask(false)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(answers, pagingData, filterData))
	}
}
