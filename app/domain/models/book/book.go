package book

import "github.com/yuonoda/bookspace/app/domain/models/user"

type Book struct {
	BookID  ID
	OwnerID user.ID
	ISBN13  ISBN13
	Title   Title
}
