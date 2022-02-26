package firestore

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yuonoda/bookspace/app/domain/models/book"
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"google.golang.org/api/iterator"
	"log"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) SaveUser(ctx context.Context, user user.User) error {
	return nil
}
func (r *Repository) LoadUser(ctx context.Context, userID user.ID) (user.User, error) {
	return user.User{}, nil
}

func (r *Repository) LoadUsers(ctx context.Context) (user.Users, error) {
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
	return user.Users{}, nil
}

func (r *Repository) SaveBook(ctx context.Context, book book.Book) error {
	return nil
}

func (r *Repository) LoadBook(ctx context.Context, bookID book.ID) (book.Book, error) {
	return book.Book{}, nil
}
