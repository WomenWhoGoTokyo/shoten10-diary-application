package diary

import (
	"context"

	"github.com/WomenWhoGoTokyo/shoten10-diary-application/domain/model"
	"github.com/WomenWhoGoTokyo/shoten10-diary-application/usecase/repository"
)

type ListDiaryOutputPort struct {
	Diaries []*model.Diary
}

type ListDiaryUseCase struct {
	diaryRepo repository.DiaryRepository
}

func NewListDiaryUseCase(dr repository.DiaryRepository) *ListDiaryUseCase {
	return &ListDiaryUseCase{dr}
}

func (du *ListDiaryUseCase) Execute(ctx context.Context) (*ListDiaryOutputPort, error) {
	diaries, err := du.diaryRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return &ListDiaryOutputPort{diaries}, nil
}
