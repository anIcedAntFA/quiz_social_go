package usermodel

import (
	"errors"
	"social_quiz/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:roles;type:ENUM('user', 'admin')"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) Mask(isAdminOrOwner bool) {
	u.GenerateUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"password" gorm:"column:password;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"-" gorm:"column:roles;type:ENUM('user','admin');default:user"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask(isAdminOrOwner bool) {
	u.GenerateUID(common.DbTypeUser)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (u *UserLogin) TableName() string {
	return User{}.TableName()
}

var (
	ErrorEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrorEmailOrPasswordInvalid",
	)
	ErrorEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrorEmailExisted",
	)
	ErrorUserDisabledOrBanned = common.NewCustomError(
		errors.New("user has been disabled or banned"),
		"user has been disabled or banned",
		"ErrorUserDisabledOrBanned",
	)
)
