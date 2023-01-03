package usermodel

import "social_quiz/common"

type UserCredential struct {
	common.SQLModel `json:",inline"`
	UserName        string `json:"user_name" gorm:"column:user_name"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	Salt            string `json:"-" gorm:"column:salt;"`
}
