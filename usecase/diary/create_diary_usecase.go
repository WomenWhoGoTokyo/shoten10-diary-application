package diary

import (
	"context"

	"github.com/WomenWhoGoTokyo/shoten10-diary-application/domain/model"
	"github.com/WomenWhoGoTokyo/shoten10-diary-application/usecase/repository"
)

type CreateDiaryInputPort struct {
	Title       string
	Description string
}

type CreateDiaryOutputPort struct {
	Diary *model.Diary
}

type CreateDiaryUseCase struct {
	diaryRepo repository.DiaryRepository
}

func NewCreateDiaryUseCase(dr repository.DiaryRepository) *CreateDiaryUseCase {
	return &CreateDiaryUseCase{dr}
}

func (du *CreateDiaryUseCase) Execute(ctx context.Context, in *CreateDiaryInputPort) (*CreateDiaryOutputPort, error) {
	diary := &model.Diary{
		Title:       in.Title,
		Description: in.Description,
	}

	diary, err := du.diaryRepo.Store(ctx, diary)
	if err != nil {
		return nil, err
	}
	return &CreateDiaryOutputPort{diary}, nil
}
