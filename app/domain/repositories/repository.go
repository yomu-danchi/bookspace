package repositories

import (
	"github.com/yuonoda/bookspace/app/domain/models/book"
	"github.com/yuonoda/bookspace/app/domain/models/user"
)

type Repository interface {
	SaveUser(user user.User) error
	RegisterBook(book book.Book) error
}
