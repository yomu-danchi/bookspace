package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/yuonoda/bookspace/app/domain/models/book"
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"google.golang.org/api/iterator"
	"log"
)

type repository struct {
	client *firestore.Client
}

func (r *repository) SaveUser(ctx context.Context, user user.User) error {
	return nil
}
func (r *repository) LoadUser(ctx context.Context, userID user.ID) (*user.User, error) {
	return nil, nil
}

func (r *repository) LoadUsers(ctx context.Context) (*user.User, error) {
	log.Println("LoadUsers")
	store := ctxlib.GetDB(ctx)
	log.Printf("store* %+v", store)
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
	log.Printf("j: %+v", j)
	return nil, nil
}

func (r *repository) SaveBook(ctx context.Context, book book.Book) error {
	return nil
}

func (r *repository) LoadBook(ctx context.Context, bookID book.ID) (*book.Book, error) {
	return nil, nil
}
