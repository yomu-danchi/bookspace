package usecase

import (
	"context"
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

func NewUseCase(r repositories.Repository) *Usecase {
	return &Usecase{
		r,
	}
}

func (u *Usecase) CreateUser(ctx context.Context, dtoUser dto.User) (dto.User, error) {
	if dtoUser.ID != "" {
		return dto.User{}, errors.Invalid(xerrors.Errorf("user id cannot exist, id: %s", dtoUser.ID))
	}

	newUserID, err := user.GenID()
	if err != nil {
		return dto.User{}, xerrors.Errorf(": %w", err)
	}
	newUserName := user.NewName(dtoUser.Name)
	newUser := user.NewUser(newUserID, newUserName)

	if err := u.repository.SaveUser(ctx, newUser); err != nil {
		return dto.User{}, xerrors.Errorf(": %w", err)
	}
	newDtoUser := dto.ToDtoUser(newUser)
	return newDtoUser, nil
}

func (u Usecase) GetUser(ctx context.Context, userIDStr string) (dto.User, error) {
	userID := user.NewID(userIDStr)
	user, err := u.repository.LoadUser(ctx, userID)
	if err != nil {
		return dto.User{}, xerrors.Errorf(":%w", err)
	}

	ownedBooks, err := u.repository.LoadBooksOwnedBy(ctx, user.ID)
	if err != nil {
		return dto.User{}, xerrors.Errorf(":%w", err)
	}
	newDtoUserWithBooks := dto.ToDtoUserWithOwendBooks(user, ownedBooks)
	return newDtoUserWithBooks, nil
}

func (u *Usecase) GetUsers(ctx context.Context) (dto.Users, error) {
	users, err := u.repository.LoadUsers(ctx)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	dtoUsers := dto.ToDtoUsers(users)
	return dtoUsers, nil
}

func (u *Usecase) RegisterBook(ctx context.Context, dtoBook dto.Book) (dto.Book, error) {
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
	newBook := book.NewBook(newBookID, ownerID, "", newISBN13, newBookTitle)

	if err := u.repository.SaveBook(ctx, newBook); err != nil {
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

func (u *Usecase) BorrowBook(ctx context.Context, bookID string, borrowerID string) error {

	book, err := u.repository.LoadBook(ctx, book.NewID(bookID))
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	user, err := u.repository.LoadUser(ctx, user.NewID(borrowerID))
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	// 複雑になったら貸し出しエンティティを設けてもいいかも
	borrowedBook := book.UpdateBorrower(user.ID)
	u.repository.SaveBook(ctx, borrowedBook)
	return nil
}

func (u *Usecase) ReturnBook() {}
