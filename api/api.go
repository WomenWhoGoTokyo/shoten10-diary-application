package api

import (
	"database/sql"

	"github.com/gorilla/mux"

	"github.com/WomenWhoGoTokyo/shoten10-diary-application/api/handler"
	"github.com/WomenWhoGoTokyo/shoten10-diary-application/interface/repository"
	"github.com/WomenWhoGoTokyo/shoten10-diary-application/usecase/diary"
)

func BuildRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	buildProjectRoutes(r, db)
	return r
}

func buildProjectRoutes(r *mux.Router, db *sql.DB) {
	dr := repository.NewDiaryRepository(db)

	r.Handle("/diary", handler.NewCreateDiaryHandler(
		diary.NewCreateDiaryUseCase(dr))).Methods("POST")
	r.Handle("/diary/{id}", handler.NewUpdateDiaryHandler(
		diary.NewUpdateDiaryUseCase(dr))).Methods("PUT")
	r.Handle("/diary/{id}/delete", handler.NewDeleteDiaryHandler(
		diary.NewDeleteDiaryUseCase(dr))).Methods("DELETE")
	r.Handle("/diary/{id}", handler.NewGetDiaryHandler(
		diary.NewGetDiaryUseCase(dr))).Methods("GET")
	r.Handle("/diaries", handler.NewListDiaryHandler(
		diary.NewListDiaryUseCase(dr))).Methods("GET")
}
