package rest

import (
	"encoding/json"
	"fmt"
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
		// TODO エラーレスポンスを返す
		log.Fatal(fmt.Errorf("failed to get users :%w", err))
	}
	log.Printf("users: %+v", users)
	j, _ := json.Marshal(users)
	w.Write(j)
}
