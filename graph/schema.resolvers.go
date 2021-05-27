package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/yuonoda/bookspace/graph/generated"
	"github.com/yuonoda/bookspace/graph/model"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	//timestamp := time.Now().Format("2006-01-02 15:04:05")

	// TODO 作成日と更新日を追加
	book := model.Book{
		OwnerID:   input.OwnerID,
		Isbn13:    input.Isbn13,
		BookTitle: input.BookTitle,
	}
	r.DB.Create(&book)

	return &book, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }