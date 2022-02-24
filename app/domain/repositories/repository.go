package repositories

import (
	"github.com/yuonoda/bookspace/app/domain/models/book"
	"github.com/yuonoda/bookspace/app/domain/models/user"
)

type Repository interface {
	SaveUser(user user.User) error
	LoadUser(userID user.ID) (*user.User, error)
	SaveBook(book book.Book) error
	LoadBook(bookID book.ID) (*book.Book, error)
}
