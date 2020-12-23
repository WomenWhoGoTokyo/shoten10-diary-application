package diary

import (
	"context"

	"github.com/WomenWhoGoTokyo/shoten10-diary-application/domain/model"
	"github.com/WomenWhoGoTokyo/shoten10-diary-application/usecase/repository"
)

type UpdateDiaryInputPort struct {
	ID          int
	Title       string
	Description string
}

type UpdateDiaryOutputPort struct {
	Diary *model.Diary
}

type UpdateDiaryUseCase struct {
	diaryRepo repository.DiaryRepository
}

func NewUpdateDiaryUseCase(dr repository.DiaryRepository) *UpdateDiaryUseCase {
	return &UpdateDiaryUseCase{dr}
}

func (du *UpdateDiaryUseCase) Execute(ctx context.Context, in *UpdateDiaryInputPort) (*UpdateDiaryOutputPort, error) {
	diary := &model.Diary{
		ID:          in.ID,
		Title:       in.Title,
		Description: in.Description,
	}

	diary, err := du.diaryRepo.Update(ctx, diary)
	if err != nil {
		return nil, err
	}
	return &UpdateDiaryOutputPort{diary}, nil
}
