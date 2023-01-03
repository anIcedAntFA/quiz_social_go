package ginquestionlike

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	questionstorage "social_quiz/module/question/storage"
	questionlikebusiness "social_quiz/module/questionlike/business"
	questionlikestorage "social_quiz/module/questionlike/storage"
)

func HandleUserDislikeQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()

		storage := questionlikestorage.NewQuestionLikeSQLStorage(db)
		storageDecrease := questionstorage.NewQuestionMySQLStorage(db)
		business := questionlikebusiness.NewUserDislikeQuestionBusiness(storage, storageDecrease)

		if err := business.DislikeQuestion(
			ctx.Request.Context(),
			requester.GetUserId(), //should struct data
			int(uid.GetLocalID()),
		); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
