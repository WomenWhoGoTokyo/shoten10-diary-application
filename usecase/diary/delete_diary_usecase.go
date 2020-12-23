package diary

import (
	"context"

	"github.com/WomenWhoGoTokyo/shoten10-diary-application/usecase/repository"
)

type DeleteDiaryInputPort struct {
	ID int
}

type DeleteDiaryUseCase struct {
	diaryRepo repository.DiaryRepository
}

func NewDeleteDiaryUseCase(dr repository.DiaryRepository) *DeleteDiaryUseCase {
	return &DeleteDiaryUseCase{dr}
}

func (du *DeleteDiaryUseCase) Execute(ctx context.Context, in *DeleteDiaryInputPort) error {
	return du.diaryRepo.Delete(ctx, in.ID)
}
