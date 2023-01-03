package usermodel

import "social_quiz/common"

type UserProfile struct {
	LastName  string        `json:"last_name" gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	Gender    string        `json:"gender" gorm:"column:phone;type:ENUM('male', 'female', 'other');"`
	Birthday  string        `json:"birthday" gorm:"column:birthday;"`
	Location  string        `json:"location" gorm:"column:location;"`
	Role      string        `json:"role" gorm:"column:roles;type:ENUM('user', 'admin');"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}
