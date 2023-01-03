package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"social_quiz/common"
	"social_quiz/components/appctx"
)

func RequiredRole(appCtx appctx.AppContext, allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false

		for _, role := range allowedRoles {
			if requester.GetRole() == role {
				hasFound = true
				break
			}
		}

		if !hasFound {
			panic(common.ErrorNoPermission(errors.New("invalid role")))
		}

		ctx.Next()
	}
}
