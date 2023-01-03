package answermodel

type Filter struct {
	Content string `json:"content" form:"content"`
	Correct string `json:"correct" form:"correct"`
	Status  []int  `json:"-"`
}
