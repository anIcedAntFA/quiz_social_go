package questionrepository

import (
	"context"
	"social_quiz/common"
	questionmodel "social_quiz/module/question/model"
)

type ListQuestionsStorage interface {
	ListQuestions(
		ctx context.Context,
		filter *questionmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]questionmodel.Question, error)
}

//type ListQuestionLikesStorage interface {
//	ListQuestionLikes(ctx context.Context, ids []int) (map[int]int, error)
//}

type listQuestionsRepository struct {
	storage ListQuestionsStorage
	//storageLike ListQuestionLikesStorage
}

func NewListQuestionsRepository(storage ListQuestionsStorage) *listQuestionsRepository {
	return &listQuestionsRepository{
		storage: storage,
		//storageLike: storageLike,
	}
}

func (biz listQuestionsRepository) ListQuestions(
	ctx context.Context,
	filter *questionmodel.Filter,
	paging *common.Paging,
) ([]questionmodel.Question, error) {
	result, err := biz.storage.ListQuestions(ctx, filter, paging, "Answers", "User")

	if err != nil {
		return nil, common.ErrorCannotListEntity(questionmodel.EntityName, err)
	}

	//ids := make([]int, len(result))
	//
	//for i := range ids {
	//	ids[i] = result[i].Id
	//}
	//
	//likeMap, err := biz.storageLike.ListQuestionLikes(ctx, ids)
	//
	//if err != nil {
	//	log.Println(err)
	//	return result, nil
	//}
	//
	//for i, value := range result {
	//	result[i].LikeCount = likeMap[value.Id]
	//}

	return result, nil
}
