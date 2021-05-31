package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"strconv"

	"github.com/yuonoda/bookspace/graph/generated"
	"github.com/yuonoda/bookspace/graph/model"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	//timestamp := time.Now().Format("2006-01-02 15:04:05")

	// TODO 作成日と更新日を追加
	userId, err := strconv.Atoi(input.UserID)
	if err != nil {
		log.Fatal(err)
	}
	book := model.Book{
		UserID:    userId,
		Isbn13:    input.Isbn13,
		BookTitle: input.BookTitle,
	}
	r.DB.Select("UserID", "Isbn13", "BookTitle").Create(&book) // フィールドを指定しない方法ある？

	return &book, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		UserName:        input.UserName,
		UserDisplayName: input.UserDisplayName,
	}
	r.DB.Select("UserName", "UserDisplayName").Create(&user)
	return &user, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	r.DB.Find(&books)
	return books, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	r.DB.Find(&users)
	log.Printf("users:%+v\n", users)
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
