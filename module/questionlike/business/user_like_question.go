package questionlikebusiness

import (
	"context"
	"log"
	"social_quiz/common"
	questionlikemodel "social_quiz/module/questionlike/model"
)

type UserLikeQuestionStorage interface {
	CreateQuestionLike(ctx context.Context, data *questionlikemodel.Like) error
	//Find
}

type IncreaseLikeCountQuestionStorage interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeQuestionBusiness struct {
	storage    UserLikeQuestionStorage
	storageInc IncreaseLikeCountQuestionStorage
}

func NewUserLikeQuestionBusiness(
	storage UserLikeQuestionStorage,
	storageInc IncreaseLikeCountQuestionStorage,
) *userLikeQuestionBusiness {
	return &userLikeQuestionBusiness{
		storage:    storage,
		storageInc: storageInc,
	}
}

func (biz *userLikeQuestionBusiness) LikeQuestion(ctx context.Context, data *questionlikemodel.Like) error {
	//check if question exist with its ID
	//and status == 0 ?

	err := biz.storage.CreateQuestionLike(ctx, data)

	if err != nil {
		return questionlikemodel.ErrorCannotLikeQuestion(err)
	}

	//side effect, flow not important
	go func() {
		defer common.AppRecover()

		if err := biz.storageInc.IncreaseLikeCount(ctx, data.QuestionId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
