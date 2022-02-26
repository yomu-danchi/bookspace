package firestore

import (
	"context"
	"github.com/yuonoda/bookspace/app/domain/models/book"
	"github.com/yuonoda/bookspace/app/domain/models/user"
)

func (r *Repository) SaveBook(ctx context.Context, book book.Book) error {
	return nil
}

func (r *Repository) LoadBook(ctx context.Context, bookID book.ID) (book.Book, error) {
	return book.Book{}, nil
}

func (r *Repository) LoadBooksOwnedBy(ctx context.Context, userID user.ID) (book.Books, error) {
	return nil, nil
}
