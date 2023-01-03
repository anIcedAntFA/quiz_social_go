package userstorage

import (
	"context"
	"social_quiz/common"
	usermodel "social_quiz/module/user/model"
)

func (s *userMySQLStorage) ListUsers(
	ctx context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	if f := filter; f != nil {
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
		if len(f.Email) > 0 {
			db = db.Where("status in (?)", f.Email)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit

	var result []usermodel.User

	if err := db.
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).
		Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return result, nil
}
