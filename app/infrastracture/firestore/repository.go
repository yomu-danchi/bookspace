package firestore

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yuonoda/bookspace/app/domain/models/book"
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"golang.org/x/xerrors"
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
	store := ctxlib.GetDB(ctx)
	iter := store.Collection("users").Documents(ctx)
	var fetched []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
		fetched = append(fetched, doc.Data())
	}

	bytes, err := json.Marshal(fetched)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse to json : %w", err)
	}
	var users user.Users
	if err := json.Unmarshal(bytes, &users); err != nil {
		return nil, xerrors.Errorf("failed to parse from json : %w", err)
	}
	return users, nil
}

func (r *Repository) SaveBook(ctx context.Context, book book.Book) error {
	return nil
}

func (r *Repository) LoadBook(ctx context.Context, bookID book.ID) (book.Book, error) {
	return book.Book{}, nil
}

func (r *Repository) LoadBooksOwnedBy(ctx context.Context, userID user.ID) (book.Books, error) {
	return nil, nil
}
