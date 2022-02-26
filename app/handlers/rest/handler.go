package rest

import (
	"encoding/json"
	"fmt"
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"github.com/yuonoda/bookspace/app/usecase"
	"google.golang.org/api/iterator"
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
	store := ctxlib.GetDB(ctx)
	iter := store.Collection("users").Documents(ctx)
	var users []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
		users = append(users, doc.Data())
	}
	j, _ := json.Marshal(users)
	w.Write(j)
}
