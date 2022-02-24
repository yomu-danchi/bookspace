package book

import "github.com/yuonoda/bookspace/app/domain/models/user"

type Book struct {
	ID      ID
	OwnerID user.ID
	ISBN13  ISBN13 // TODO nullableにする
	Title   Title
}

func NewBook(ID ID, OwnerID user.ID, isbn13 ISBN13, title Title) Book {
	return Book{
		ID:      ID,
		OwnerID: OwnerID,
		ISBN13:  isbn13,
		Title:   title,
	}
}
