package rest

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yuonoda/bookspace/app/usecase"
	"github.com/yuonoda/bookspace/app/usecase/dto"
	"golang.org/x/xerrors"
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

func (h handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var dtoUser dto.User
	// TODO JSON が間違っていたときのバリデーション
	err := json.NewDecoder(r.Body).Decode(&dtoUser)
	if err != nil {
		render.Status(r, 400)
		return
	}
	user, err := h.usecase.CreateUser(ctx, dtoUser)
	if err != nil {
		log.Print(xerrors.Errorf(":%w", err))
		render.Status(r, 404)
	}
	render.JSON(w, r, user)
}

func (h handler) RegisterBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var dtoBook dto.Book
	// TODO JSON が間違っていたときのバリデーション
	err := json.NewDecoder(r.Body).Decode(&dtoBook)
	if err != nil {
		render.Status(r, 400)
		return
	}
	user, err := h.usecase.RegisterBook(ctx, dtoBook)
	if err != nil {
		log.Print(xerrors.Errorf(":%w", err))
		render.Status(r, 400)
	}
	render.JSON(w, r, user)
}

func (h handler) ReturnBook(w http.ResponseWriter, r *http.Request) {}
func (h handler) BorrowBook(w http.ResponseWriter, r *http.Request) {}
