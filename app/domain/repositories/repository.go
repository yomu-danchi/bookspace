package repositories

import (
	"context"
	"github.com/yuonoda/bookspace/app/domain/models/book"
	"github.com/yuonoda/bookspace/app/domain/models/user"
)

type Repository interface {
	// TODO not foundのときの対応
	SaveUser(ctx context.Context, user user.User) error
	LoadUser(ctx context.Context, userID user.ID) (user.User, error)
	LoadUsers(ctx context.Context) (user.Users, error)
	SaveBook(ctx context.Context, book book.Book) error
	LoadBook(ctx context.Context, bookID book.ID) (book.Book, error)
	LoadBooksOwnedBy(ctx context.Context, userID user.ID) (book.Books, error)
}
