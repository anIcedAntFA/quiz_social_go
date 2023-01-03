package ginuser

import (
	"net/http"
	"social_quiz/common"
	"social_quiz/components/appctx"

	"github.com/gin-gonic/gin"
)

func HandleGetProfileUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
