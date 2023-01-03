package questionlikebusiness

import (
	"context"
	"log"
	"social_quiz/common"
	questionlikemodel "social_quiz/module/questionlike/model"
)

type UserDislikeQuestionStorage interface {
	DeleteQuestionLike(ctx context.Context, userId, questionId int) error
	//Find
}

type DecreaseLikeCountQuestionStorage interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userDislikeQuestionBusiness struct {
	storage    UserDislikeQuestionStorage
	storageDec DecreaseLikeCountQuestionStorage
}

func NewUserDislikeQuestionBusiness(
	storage UserDislikeQuestionStorage,
	storageDec DecreaseLikeCountQuestionStorage,
) *userDislikeQuestionBusiness {
	return &userDislikeQuestionBusiness{
		storage:    storage,
		storageDec: storageDec,
	}
}

func (biz *userDislikeQuestionBusiness) DislikeQuestion(ctx context.Context, userId, questionId int) error {
	//check if question exist with its ID
	//and status == 0 ?

	err := biz.storage.DeleteQuestionLike(ctx, userId, questionId)

	if err != nil {
		return questionlikemodel.ErrorCannotDislikeQuestion(err)
	}

	go func() {
		defer common.AppRecover()

		if err := biz.storageDec.DecreaseLikeCount(ctx, questionId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
