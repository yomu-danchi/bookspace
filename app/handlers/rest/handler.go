package rest

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/yuonoda/bookspace/app/handlers/rest/middleware"
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
)

func Serve() {
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(middleware.SetDB)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("users: %+v", users)
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":8000", r)
}
