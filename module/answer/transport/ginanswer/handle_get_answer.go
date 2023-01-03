package ginanswer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	answerbusiness "social_quiz/module/answer/business"
	answerstorage "social_quiz/module/answer/storage"
)

func HandleGetAnswer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := answerstorage.NewAnswerMySQLStorage(db)
		business := answerbusiness.NewFindAnswerBusiness(storage)

		answer, err := business.FindAnswer(
			ctx.Request.Context(),
			map[string]interface{}{"id": int(uid.GetLocalID()), "status": 1},
		)

		if err != nil {
			panic(err)
		}

		answer.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(answer))
	}
}
