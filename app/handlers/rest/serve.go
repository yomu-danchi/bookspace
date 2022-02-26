package rest

import (
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/yuonoda/bookspace/app/handlers/rest/middleware"
	"github.com/yuonoda/bookspace/app/handlers/rest/oapi"
	"github.com/yuonoda/bookspace/app/infrastracture/firestore"
	"github.com/yuonoda/bookspace/app/usecase"
	"net/http"
)

func Serve() {
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(middleware.SetDB)

	// TODO wireを使ってDI
	repo := firestore.NewRepository()
	u := usecase.NewUseCase(repo)
	h := NewHandler(u)

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Welcome to BookSpace API"))
	})
	oapi.HandlerFromMux(h, r)
	http.ListenAndServe(":8000", r)
}
