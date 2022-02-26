package dto

import "github.com/yuonoda/bookspace/app/domain/models/book"

type Books []Book

func ToDtoBooks(domainBooks book.Books) Books {
	books := make(Books, 0, len(domainBooks))
	for _, domainBook := range domainBooks {
		books = append(books, ToDtoBook(domainBook))
	}
	return books
}
