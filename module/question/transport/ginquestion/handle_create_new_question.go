package ginquestion

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	questionbusiness "social_quiz/module/question/business"
	questionmodel "social_quiz/module/question/model"
	questionstorage "social_quiz/module/question/storage"
)

func HandleCreateNewQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		var newData questionmodel.QuestionCreate

		if err := ctx.ShouldBind(&newData); err != nil {
			panic(err)
		}

		newData.UserId = requester.GetUserId()

		storage := questionstorage.NewQuestionMySQLStorage(db)
		business := questionbusiness.NewCreateQuestionBusiness(storage)

		if err := business.CreateQuestion(ctx.Request.Context(), &newData); err != nil {
			panic(err)
		}

		newData.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(newData.FakeId.String()))
	}
}
