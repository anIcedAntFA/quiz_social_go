package userstorage

import (
	"context"
	"social_quiz/common"
	usermodel "social_quiz/module/user/model"

	"gorm.io/gorm"
)

func (s *userMySQLStorage) FindUser(
	ctx context.Context,
	condition map[string]interface{},
	moreInfo ...string,
) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}

		return nil, common.ErrorDB(err)
	}

	return &user, nil
}
