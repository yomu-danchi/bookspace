package rest

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yuonoda/bookspace/app/usecase"
	"log"
	"net/http"
)

type handler struct {
	usecase usecase.Usecase
}

func NewHandler(u *usecase.Usecase) *handler {
	return &handler{
		*u,
	}
}

func (h handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := h.usecase.GetUsers(ctx)
	if err != nil {
		log.Print(fmt.Errorf("failed to get users :%w", err))
		render.Status(r, 404)
	}
	render.JSON(w, r, users)
}

func (h handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := chi.URLParam(r, "userID")
	user, err := h.usecase.GetUser(ctx, userID)
	if err != nil {
		log.Print(fmt.Errorf("failed to get users :%w", err))
		render.Status(r, 404)
	}
	render.JSON(w, r, user)
}
