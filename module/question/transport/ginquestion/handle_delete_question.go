package ginquestion

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	questionbusiness "social_quiz/module/question/business"
	questionstorage "social_quiz/module/question/storage"
)

func HandleDeleteQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			//because err is golang error
			panic(common.ErrorInvalidRequest(err))
		}

		storage := questionstorage.NewQuestionMySQLStorage(db)
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		business := questionbusiness.NewDeleteQuestionBusiness(storage, requester)

		if err := business.DeleteQuestion(ctx.Request.Context(), int(uid.GetLocalID())); err != nil {
			//because business layer has been handle this error
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
