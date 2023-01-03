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

func HandleUpdateQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrorInvalidRequest(err))
			return
		}

		var updatedData questionmodel.QuestionUpdate

		if err := ctx.ShouldBind(&updatedData); err != nil {
			// panic(err)
			ctx.JSON(http.StatusBadRequest, common.ErrorInvalidRequest(err))

			return
		}

		storage := questionstorage.NewQuestionMySQLStorage(db)
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		business := questionbusiness.NewUpdateQuestionBusiness(storage, requester)

		if err := business.UpdateQuestion(
			ctx.Request.Context(),
			int(uid.GetLocalID()),
			&updatedData,
		); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				common.ErrorCannotUpdateEntity(questionmodel.EntityName, err),
			)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
