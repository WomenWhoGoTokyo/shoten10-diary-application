package presenter

import (
	"github.com/WomenWhoGoTokyo/shoten10-diary-application/domain/model"
	"github.com/WomenWhoGoTokyo/shoten10-diary-application/usecase/diary"
)

type simpleDiaryView struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func mapDiaryToSimpleView(diary *model.Diary) *simpleDiaryView {
	return &simpleDiaryView{
		ID:          diary.ID,
		Title:       diary.Title,
		Description: diary.Description,
	}
}

type createDiaryResponse struct {
	Diary *simpleDiaryView `json:"diary"`
}

func NewCreateDiaryPresenter(output *diary.CreateDiaryOutputPort) *createDiaryResponse {
	return &createDiaryResponse{mapDiaryToSimpleView(output.Diary)}
}

type updateDiaryResponse struct {
	Diary *simpleDiaryView `json:"diary"`
}

func NewUpdateDiaryPresenter(output *diary.UpdateDiaryOutputPort) *updateDiaryResponse {
	return &updateDiaryResponse{mapDiaryToSimpleView(output.Diary)}
}

type getDiaryResponse struct {
	Diary *simpleDiaryView `json:"diary"`
}

func NewGetDiaryPresenter(output *diary.GetDiaryOutputPort) *getDiaryResponse {
	return &getDiaryResponse{mapDiaryToSimpleView(output.Diary)}
}

type listDiaryResponse struct {
	Diaries []*simpleDiaryView `json:"diaries"`
}

func NewListDiaryPresenter(output *diary.ListDiaryOutputPort) *listDiaryResponse {
	views := make([]*simpleDiaryView, len(output.Diaries))

	for i, row := range output.Diaries {
		views[i] = mapDiaryToSimpleView(row)
	}
	return &listDiaryResponse{views}
}
