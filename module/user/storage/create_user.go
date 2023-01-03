package userstorage

import (
	"context"
	"social_quiz/common"
	usermodel "social_quiz/module/user/model"
)

func (s *userMySQLStorage) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	//if u implement many action => open transaction & commit
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		//if err, call rollback before, if not, connection will be stuck here
		//=> too many connection to DB => crash DB
		db.Rollback()

		return common.ErrorDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()

		return common.ErrorDB(err)
	}

	return nil
}
