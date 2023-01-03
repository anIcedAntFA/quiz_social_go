package common

import "time"

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"updated_at;"`
}

func (s *SQLModel) GenerateUID(dbType int) {
	uid := NewUID(uint32(s.Id), dbType, 1)

	s.FakeId = &uid //pointer for fakeid =nil if empty
}