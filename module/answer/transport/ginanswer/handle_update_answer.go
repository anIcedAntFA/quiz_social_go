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

func HandleUpdateAnswer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		var updatedAnswer answermodel.AnswerUpdate

		if err := ctx.ShouldBind(&updatedAnswer); err != nil {
			panic(err)
		}

		storage := answerstorage.NewAnswerMySQLStorage(db)
		business := answerbusiness.NewUpdateAnswerBusiness(storage)

		if err := business.UpdateAnswer(
			ctx.Request.Context(),
			int(uid.GetLocalID()),
			&updatedAnswer,
		); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
