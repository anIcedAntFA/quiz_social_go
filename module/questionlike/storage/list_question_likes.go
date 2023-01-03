package questionlikestorage

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcd/btcutil/base58"
	"social_quiz/common"
	questionlikemodel "social_quiz/module/questionlike/model"
	"time"
)

func (s *questionLikeMySQLStorage) ListQuestionLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type mySQLData struct {
		QuestionId int `gorm:"column:question_id;"`
		LikeCount  int `gorm:"column:count;"`
	}

	var listLikes []mySQLData

	if err := s.db.Table(questionlikemodel.Like{}.TableName()).
		Select("question_id, count(question_id) as count").
		Where("question_id in (?)", ids).
		Group("question_id").
		Find(&listLikes).
		Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	for _, item := range listLikes {
		result[item.QuestionId] = item.LikeCount
	}

	return result, nil
}

func (s *questionLikeMySQLStorage) ListUsersLikeQuestion(
	ctx context.Context,
	conditions map[string]interface{},
	filter *questionlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	db := s.db

	db = db.Table(questionlikemodel.Like{}.TableName()).Where(conditions)

	if f := filter; f != nil {
		if f.QuestionId > 0 {
			db = db.Where("question_id = ?", f.QuestionId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		createdTime, err := time.Parse(common.TimeLayout, string(base58.Decode(v)))

		if err != nil {
			return nil, common.ErrorDB(err)
		}

		db = db.Where("created_at < ?", createdTime.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	var result []questionlikemodel.Like

	if err := db.Limit(paging.Limit).Order("created_at desc").Find(&result).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	for i, v := range result {
		result[i].User.CreatedAt = v.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User

		if i == len(result)-1 {
			cursorString := base58.Encode([]byte(fmt.Sprintf("%v", v.CreatedAt.Format(common.TimeLayout))))
			paging.NextCursor = cursorString
		}
	}

	return users, nil
}
