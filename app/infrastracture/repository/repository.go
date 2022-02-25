package repository

import (
	"cloud.google.com/go/firestore"
	"github.com/yuonoda/bookspace/app/domain/models/book"
	"github.com/yuonoda/bookspace/app/domain/models/user"
)

type repository struct {
	client *firestore.Client
}

func (r *repository) SaveUser(user user.User) error {
	return nil
}
func (r *repository) LoadUser(userID user.ID) (*user.User, error) {
	return nil, nil
}

func (r *repository) SaveBook(book book.Book) error {
	return nil
}

func (r *repository) LoadBook(bookID book.ID) (*book.Book, error) {
	return nil, nil
}
