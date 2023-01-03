package ginanswer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	answerbusiness "social_quiz/module/answer/business"
	answerstorage "social_quiz/module/answer/storage"
)

func HandleDeleteAnswer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := answerstorage.NewAnswerMySQLStorage(db)
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		business := answerbusiness.NewDeleteAnswerBusiness(storage, requester)

		if err := business.DeleteAnswer(ctx.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
