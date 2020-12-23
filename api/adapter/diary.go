package adapter

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/WomenWhoGoTokyo/shoten10-diary-application/usecase/diary"
)

type createDiaryRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewCreateDiaryInputPortFromRequest(r *http.Request) (*diary.CreateDiaryInputPort, error) {
	var input createDiaryRequestBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, err
	}

	return &diary.CreateDiaryInputPort{
		Title:       input.Title,
		Description: input.Description,
	}, nil
}

type updateDiaryRequestBody struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewUpdateDiaryInputPortFromRequest(r *http.Request) (*diary.UpdateDiaryInputPort, error) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	var input updateDiaryRequestBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, err
	}

	return &diary.UpdateDiaryInputPort{
		ID:          ID,
		Title:       input.Title,
		Description: input.Description,
	}, nil
}

func NewDeleteDiaryInputPortFromRequest(r *http.Request) (*diary.DeleteDiaryInputPort, error) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	return &diary.DeleteDiaryInputPort{
		ID: ID,
	}, nil
}

func NewGetDiaryInputPortFromRequest(r *http.Request) (*diary.GetDiaryInputPort, error) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	return &diary.GetDiaryInputPort{
		ID: ID,
	}, nil
}
