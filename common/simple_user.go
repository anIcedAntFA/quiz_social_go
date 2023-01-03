package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	Role      string `json:"role" gorm:"column:roles;type:ENUM('user', 'admin')"`
	Avatar    *Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (s *SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdminOrOwner bool) {
	u.GenerateUID(DbTypeUser)
}
