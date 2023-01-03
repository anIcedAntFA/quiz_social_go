package ginanswer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	"social_quiz/module/answer/business"
	answermodel "social_quiz/module/answer/model"
	answerstorage "social_quiz/module/answer/storage"
)

func HandleCreateNewAnswer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		var newAnswer answermodel.AnswerCreate

		if err := ctx.ShouldBind(&newAnswer); err != nil {
			panic(err)
		}

		newAnswer.QuestionId = int(uid.GetLocalID())

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		newAnswer.UserId = requester.GetUserId()

		db := appCtx.GetMainDBConnection()

		storage := answerstorage.NewAnswerMySQLStorage(db)
		business := answerbusiness.NewCreateAnswerBusiness(storage)

		if err := business.CreateAnswer(ctx, &newAnswer); err != nil {
			panic(err)
		}

		newAnswer.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(newAnswer.FakeId.String()))
	}
}
