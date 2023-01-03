package questionlikemodel

type Filter struct {
	QuestionId int `json:"-" form:"question_id"` // any user liked question ?
	UserId     int `json:"-" form:"user_id"`     // 1 user like any question ?
}
