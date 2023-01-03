package questionstorage

import (
	"context"
	"social_quiz/common"
	questionmodel "social_quiz/module/question/model"
)

func (s *questionMySQLStorage) ListQuestions(
	ctx context.Context,
	filter *questionmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]questionmodel.Question, error) {
	//requester := ctxs.MustGet(common.CurrentUser).(common.Requester)
	db := s.db.Table(questionmodel.Question{}.TableName())

	if f := filter; f != nil {
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
		if len(f.Type) > 0 {
			db = db.Where("type in (?)", f.Type)
		}
		if len(f.Category) > 0 {
			db = db.Where("category in (?)", f.Category)
		}
		if len(f.Level) > 0 {
			db = db.Where("level in (?)", f.Level)
		}
		if f.Score > 0 {
			db = db.Where("level in (?)", f.Score)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	db = db.
		Preload("User", "status = 1").
		Preload("Answers", "status = 1")
	//for _, v := range moreKeys {
	//	db = db.Preload(v)
	//}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrorDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	var result []questionmodel.Question

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).
		Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
