package main

import (
	"social_quiz/components/appctx"
	"social_quiz/middleware"
	"social_quiz/module/answer/transport/ginanswer"
	"social_quiz/module/question/transport/ginquestion"
	"social_quiz/module/questionlike/transport/ginquestionlike"
	"social_quiz/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func setupMainRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	questions := v1.Group("/questions")
	questions.POST("", ginquestion.HandleCreateNewQuestion(appCtx))
	questions.GET("/:id", ginquestion.HandleGetQuestion(appCtx))
	questions.GET("", ginquestion.HandleGetListQuestions(appCtx))
	questions.PATCH("/:id", ginquestion.HandleUpdateQuestion(appCtx))
	questions.DELETE("/:id", ginquestion.HandleDeleteQuestion(appCtx))

	answers := v1.Group("/answers", middleware.RequiredAuthorization(appCtx))
	answers.POST("/:id", ginanswer.HandleCreateNewAnswer(appCtx))
	answers.GET("/:id", ginanswer.HandleGetAnswer(appCtx))
	answers.GET("", ginanswer.HandleGetListAnswers(appCtx))
	answers.PATCH("/:id", ginanswer.HandleUpdateAnswer(appCtx))
	answers.DELETE("/:id", ginanswer.HandleDeleteAnswer(appCtx))

	auth := v1.Group("/auth")
	auth.POST("/register", ginuser.HandleRegisterUser(appCtx))
	auth.POST("/authenticate", ginuser.HandleLoginUser(appCtx))

	user := v1.Group("user")
	user.GET(
		"/profile",
		middleware.RequiredAuthorization(appCtx),
		ginuser.HandleGetProfileUser(appCtx),
	)

	likedUsers := questions.Group("/:id/liked-users")
	likedUsers.POST("", ginquestionlike.HandleUserLikeQuestion(appCtx))
	likedUsers.DELETE("", ginquestionlike.HandleUserDislikeQuestion(appCtx))
	likedUsers.GET("", ginquestionlike.HandleGetListUsersLikeQuestion(appCtx))
}
