package ginquestionlike

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"
	questionstorage "social_quiz/module/question/storage"
	questionlikebusiness "social_quiz/module/questionlike/business"
	questionlikemodel "social_quiz/module/questionlike/model"
	questionlikestorage "social_quiz/module/questionlike/storage"
)

func HandleUserLikeQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		result := questionlikemodel.Like{
			QuestionId: int(uid.GetLocalID()),
			UserId:     requester.GetUserId(),
		}

		db := appCtx.GetMainDBConnection()

		storage := questionlikestorage.NewQuestionLikeSQLStorage(db)
		storageIncrease := questionstorage.NewQuestionMySQLStorage(db)
		business := questionlikebusiness.NewUserLikeQuestionBusiness(storage, storageIncrease)

		if err := business.LikeQuestion(ctx.Request.Context(), &result); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
