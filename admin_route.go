package main

import (
	"github.com/gin-gonic/gin"
	"social_quiz/components/appctx"
	"social_quiz/middleware"
	"social_quiz/module/user/transport/ginuser"
)

func setupAdminRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group(
		"/admin",
		middleware.RequiredAuthorization(appCtx),
		middleware.RequiredRole(appCtx, "admin", "mod"),
	)
	admin.GET("/users", ginuser.HandleGetListUsers(appCtx))

	{
		admin.GET("/profile", ginuser.HandleGetProfileUser(appCtx))
	}
}
