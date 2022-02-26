package dto

import "github.com/yuonoda/bookspace/app/domain/models/book"

type Book struct {
	ID         string
	OwnerID    string
	BorrowerID string
	ISBN13     string
	Title      string
}

func ToDtoBook(domainBook book.Book) Book {
	return Book{
		ID:         domainBook.ID.String(),
		OwnerID:    domainBook.OwnerID.String(),
		BorrowerID: domainBook.BorrowerID.String(),
		ISBN13:     domainBook.ISBN13.String(),
		Title:      domainBook.Title.String(),
	}
}
