package usermodel

type Filter struct {
	Email     string `json:"email" gorm:"column:email;"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	Phone     string `json:"phone" gorm:"column:phone;"`
	Role      string `json:"role" gorm:"column:roles;type:ENUM('user', 'admin')"`
	Status    []int  `json:"-"`
}
