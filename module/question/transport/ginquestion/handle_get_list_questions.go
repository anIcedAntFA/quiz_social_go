package ginquestion

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	questionbusiness "social_quiz/module/question/business"
	questionmodel "social_quiz/module/question/model"
	questionrepository "social_quiz/module/question/repository"
	questionstorage "social_quiz/module/question/storage"
)

func HandleGetListQuestions(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := ctx.ShouldBind(&pagingData); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filterData questionmodel.Filter

		if err := ctx.ShouldBind(&filterData); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		if requester.GetRole() == "admin" {
			filterData.Status = []int{1}
		}

		storage := questionstorage.NewQuestionMySQLStorage(db)
		//storageLike := questionlikestorage.NewQuestionLikeSQLStorage(db)
		repository := questionrepository.NewListQuestionsRepository(storage)
		business := questionbusiness.NewListQuestionsBusiness(repository)

		var questions []questionmodel.Question

		questions, err := business.ListQuestions(ctx.Request.Context(), &filterData, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range questions {
			questions[i].Mask(false)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(questions, pagingData, filterData))
	}
}
