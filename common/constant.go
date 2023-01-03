package common

import "log"

const (
	DbTypeQuestion = 1
	DbTypeAnswer   = 2
	DbTypeUser     = 4
)

const CurrentUser = "user"

const TimeLayout = "2006-01-02T15:04:05.999999"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

type QuestionRequest interface {
	GetQuestionId() int
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
