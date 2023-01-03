package questionlikestorage

import "gorm.io/gorm"

type questionLikeMySQLStorage struct {
	db *gorm.DB
}

func NewQuestionLikeSQLStorage(db *gorm.DB) *questionLikeMySQLStorage {
	return &questionLikeMySQLStorage{db: db}
}
