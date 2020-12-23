package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/WomenWhoGoTokyo/shoten10-diary-application/api/adapter"
	"github.com/WomenWhoGoTokyo/shoten10-diary-application/api/presenter"
	"github.com/WomenWhoGoTokyo/shoten10-diary-application/usecase/diary"
)

func NewCreateDiaryHandler(du *diary.CreateDiaryUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := adapter.NewCreateDiaryInputPortFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		output, err := du.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewCreateDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewUpdateDiaryHandler(du *diary.UpdateDiaryUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := adapter.NewUpdateDiaryInputPortFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		output, err := du.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewUpdateDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewDeleteDiaryHandler(du *diary.DeleteDiaryUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := adapter.NewDeleteDiaryInputPortFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := du.Execute(r.Context(), input); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO: Delete Response
	})
}

func NewGetDiaryHandler(du *diary.GetDiaryUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := adapter.NewGetDiaryInputPortFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		output, err := du.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewGetDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewListDiaryHandler(du *diary.ListDiaryUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		output, err := du.Execute(r.Context())
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewListDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
