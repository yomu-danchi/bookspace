package book

import "github.com/yuonoda/bookspace/app/domain/models/user"

type Book struct {
	ID         ID
	OwnerID    user.ID
	BorrowerID user.ID
	ISBN13     ISBN13 // TODO nullableにする
	Title      Title
}

func NewBook(ID ID, ownerID user.ID, borrowerID user.ID, isbn13 ISBN13, title Title) Book {
	return Book{
		ID:         ID,
		OwnerID:    ownerID,
		BorrowerID: borrowerID,
		ISBN13:     isbn13,
		Title:      title,
	}
}

func (b Book) UpdateBorrower(borrowerID user.ID) Book {
	return Book{
		ID:         b.ID,
		OwnerID:    b.OwnerID,
		BorrowerID: borrowerID,
		ISBN13:     b.ISBN13,
		Title:      b.Title,
	}
}
