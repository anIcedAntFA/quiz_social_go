package questionmodel

type Filter struct {
	Category string `json:"category" form:"category"`
	Type     string `json:"type" form:"type"`
	Level    string `json:"level" form:"level"`
	Score    int    `json:"score" form:"score"`
	Status   []int  `json:"-"`
}
