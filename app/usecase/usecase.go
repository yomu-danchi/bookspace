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
		return dto.User{}, xerrors.Errorf(": %w", err)
	}
	newUserName := user.NewName(dtoUser.Name)
	newUser := user.NewUser(newUserID, newUserName)

	if err := u.repository.SaveUser(newUser); err != nil {
		return dto.User{}, xerrors.Errorf(": %w", err)
	}
	newDtoUser := dto.ToDtoUser(newUser)
	return newDtoUser, nil
}

func (u *Usecase) GetUsers() (dto.Users, error) {
	users, err := u.repository.LoadUsers()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	dtoUsers := dto.ToDtoUsers(users)
	return dtoUsers, nil
}

func (u *Usecase) RegisterBook(dtoBook dto.Book) (dto.Book, error) {
	if dtoBook.ID != "" {
		return dto.Book{}, errors.Invalid(xerrors.Errorf("book id cannot exist, id: %s", dtoBook.ID))
	}

	// TODO OwnerIDの存在確認
	ownerID := user.NewID(dtoBook.OwnerID)

	newBookID, err := book.GenID()
	if err != nil {
		return dto.Book{}, xerrors.Errorf(": %w", err)
	}
	newBookTitle := book.NewTitle(dtoBook.Title)
	newISBN13 := book.NewISBN13(dtoBook.ISBN13) // TODO ISBNの代わりに画像を登録させたい
	newBook := book.NewBook(newBookID, ownerID, nil, newISBN13, newBookTitle)

	if err := u.repository.SaveBook(newBook); err != nil {
		return dto.Book{}, xerrors.Errorf(": %w", err)
	}

	registeredBook := dto.Book{
		ID:      newBookID.String(),
		OwnerID: ownerID.String(),
		ISBN13:  newISBN13.String(),
		Title:   newBookTitle.String(),
	}

	return registeredBook, nil
}

func (u *Usecase) BorrowBook(bookID string, borrowerID string) error {

	book, err := u.repository.LoadBook(book.NewID(bookID))
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if book == nil {
		return errors.NotFound(xerrors.Errorf("book not found"))
	}

	user, err := u.repository.LoadUser(user.NewID(borrowerID))
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if book == nil {
		return errors.NotFound(xerrors.Errorf("user not found"))
	}

	// 複雑になったら貸し出しエンティティを設けてもいいかも
	borrowedBook := book.UpdateBorrower(&user.ID)
	u.repository.SaveBook(borrowedBook)
	return nil
}

func (u *Usecase) ReturnBook() {}
