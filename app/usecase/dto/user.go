package dto

import (
	"github.com/yuonoda/bookspace/app/domain/models/book"
	"github.com/yuonoda/bookspace/app/domain/models/user"
)

type User struct {
	ID         string
	Name       string
	OwnedBooks Books
}

func ToDtoUser(domainUser user.User) User {
	return User{
		ID:         domainUser.ID.String(),
		Name:       domainUser.Name.String(),
		OwnedBooks: nil,
	}
}

func ToDtoUserWithOwendBooks(domainUser user.User, owedBooks book.Books) User {
	return User{
		ID:         domainUser.ID.String(),
		Name:       domainUser.Name.String(),
		OwnedBooks: ToDtoBooks(owedBooks),
	}
}
