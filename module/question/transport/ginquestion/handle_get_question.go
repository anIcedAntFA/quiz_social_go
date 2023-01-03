package ginquestion

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	questionbusiness "social_quiz/module/question/business"
	questionstorage "social_quiz/module/question/storage"
)

func HandleGetQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := questionstorage.NewQuestionMySQLStorage(db)
		business := questionbusiness.NewFindQuestionBusiness(storage)

		question, err := business.FindQuestion(
			ctx.Request.Context(),
			map[string]interface{}{"id": int(uid.GetLocalID()), "status": 1},
		)

		if err != nil {
			panic(err)
		}

		question.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(question))
	}
}
