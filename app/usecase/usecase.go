package usecase

import (
	"github.com/yuonoda/bookspace/app/domain/models/book"
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/domain/repositories"
	"github.com/yuonoda/bookspace/app/errors"
	"github.com/yuonoda/bookspace/app/usecase/dto"
	"golang.org/x/xerrors"
)

type Usecase struct {
	repository repositories.Repository
}

func (u *Usecase) CreateUser(dtoUser dto.User) (dto.User, error) {
	if dtoUser.ID != "" {
		return dto.User{}, errors.Invalid(xerrors.Errorf("user id cannot exist, id: %s", dtoUser.ID))
	}

	newUserID, err := user.GenID()
	if err != nil {
		return dto.User{}, err
	}
	newUserName := user.NewName(dtoUser.Name)
	newUser := user.NewUser(newUserID, newUserName)

	if err := u.repository.SaveUser(newUser); err != nil {
		return dto.User{}, err
	}

	createdUser := dto.User{
		ID:   newUserID.String(),
		Name: newUserName.String(),
	}
	return createdUser, nil
}

func (u *Usecase) RegisterBook(dtoBook dto.Book) (dto.Book, error) {
	if dtoBook.ID != "" {
		return dto.Book{}, errors.Invalid(xerrors.Errorf("book id cannot exist, id: %s", dtoBook.ID))
	}

	// TODO OwnerIDの存在確認
	ownerID := user.NewID(dtoBook.OwnerID)

	newBookID, err := book.GenID()
	if err != nil {
		return dto.Book{}, err
	}
	newBookTitle := book.NewTitle(dtoBook.Title)
	newISBN13 := book.NewISBN13(dtoBook.ISBN13)
	newBook := book.NewBook(newBookID, ownerID, newISBN13, newBookTitle)

	if err := u.repository.RegisterBook(newBook); err != nil {
		return dto.Book{}, err
	}

	registeredBook := dto.Book{
		ID:      newBookID.String(),
		OwnerID: ownerID.String(),
		ISBN13:  newISBN13.String(),
		Title:   newBookTitle.String(),
	}

	return registeredBook, nil
}
