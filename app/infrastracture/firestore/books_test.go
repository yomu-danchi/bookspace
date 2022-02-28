package firestore

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/yuonoda/bookspace/app/domain/models/book"
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/errors"
	"github.com/yuonoda/bookspace/app/errors/codes"
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"log"
	"testing"
)

func TestRepository_LoadBooksOwnedBy(t *testing.T) {
	type args struct {
		ctx     context.Context
		ownerID user.ID
	}
	tests := []struct {
		name      string
		fixtures  func(ctx context.Context)
		args      args
		want      book.Books
		wantError codes.Code
	}{
		{
			name: "success",
			fixtures: func(ctx context.Context) {
				collection := testStore.Collection(BooksCollectionName)
				deleteCollection(ctx, testStore, collection, 100)
				_, err := collection.Doc("book1_ID").Set(ctx, map[string]interface{}{
					"ID":         "book1_ID",
					"OwnerID":    "user1_ID",
					"BorrowerID": "",
					"ISBN13":     "",
					"Title":      "book1_title",
				})
				if err != nil {
					log.Fatal(err)
				}
				_, err = collection.Doc("book2_ID").Set(ctx, map[string]interface{}{
					"ID":         "book2_ID",
					"OwnerID":    "user1_ID",
					"BorrowerID": "",
					"ISBN13":     "",
					"Title":      "book2_title",
				})
				if err != nil {
					log.Fatal(err)
				}
				_, err = collection.Doc("book3_ID").Set(ctx, map[string]interface{}{
					"ID":         "book3_ID",
					"OwnerID":    "user2_ID",
					"BorrowerID": "",
					"ISBN13":     "",
					"Title":      "book3_title",
				})
				if err != nil {
					log.Fatal(err)
				}
			},
			args: args{
				ctx:     context.WithValue(context.Background(), ctxlib.DBContextKey, testStore),
				ownerID: "user1_ID",
			},
			want: book.Books{
				{
					ID:         "book1_ID",
					OwnerID:    "user1_ID",
					BorrowerID: "",
					ISBN13:     "",
					Title:      "book1_title",
				},
				{
					ID:         "book2_ID",
					OwnerID:    "user1_ID",
					BorrowerID: "",
					ISBN13:     "",
					Title:      "book2_title",
				},
			},
			wantError: codes.OK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{}
			ctx := tt.args.ctx
			tt.fixtures(ctx)
			got, err := r.LoadBooksOwnedBy(ctx, tt.args.ownerID)
			if diff := cmp.Diff(errors.Code(err), tt.wantError); diff != "" {
				t.Errorf(diff)
				t.Log(err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}

}
func TestRepository_LoadBook(t *testing.T) {
	type args struct {
		ctx    context.Context
		bookID book.ID
	}
	tests := []struct {
		name      string
		fixtures  func(ctx context.Context)
		args      args
		want      book.Book
		wantError codes.Code
	}{
		{
			name: "success",
			fixtures: func(ctx context.Context) {
				collection := testStore.Collection(BooksCollectionName)
				deleteCollection(ctx, testStore, collection, 100)
				_, err := collection.Doc("book1_ID").Set(ctx, map[string]interface{}{
					"ID":         "book1_ID",
					"OwnerID":    "user1_ID",
					"BorrowerID": "",
					"ISBN13":     "",
					"Title":      "book1_title",
				})
				if err != nil {
					log.Fatal(err)
				}
				_, err = collection.Doc("book2_ID").Set(ctx, map[string]interface{}{
					"ID":         "book2_ID",
					"OwnerID":    "user1_ID",
					"BorrowerID": "",
					"ISBN13":     "",
					"Title":      "book2_title",
				})
				if err != nil {
					log.Fatal(err)
				}
			},
			args: args{
				ctx:    context.WithValue(context.Background(), ctxlib.DBContextKey, testStore),
				bookID: "book2_ID",
			},
			want: book.Book{
				ID:         "book2_ID",
				OwnerID:    "user1_ID",
				BorrowerID: "",
				ISBN13:     "",
				Title:      "book2_title",
			},
			wantError: codes.OK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{}
			ctx := tt.args.ctx
			tt.fixtures(ctx)
			got, err := r.LoadBook(ctx, tt.args.bookID)
			if diff := cmp.Diff(errors.Code(err), tt.wantError); diff != "" {
				t.Errorf(diff)
				t.Log(err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}

}
func TestRepository_SaveBook(t *testing.T) {
	type args struct {
		ctx  context.Context
		book book.Book
	}
	tests := []struct {
		name      string
		args      args
		fixture   func(ctx context.Context)
		want      book.Book
		wantError codes.Code
	}{
		{
			name: "success",
			fixture: func(ctx context.Context) {
				collection := testStore.Collection(BooksCollectionName)
				deleteCollection(ctx, testStore, collection, 100)
			},
			args: args{
				ctx: context.WithValue(context.Background(), ctxlib.DBContextKey, testStore),
				book: book.Book{
					ID:         "book1_id",
					OwnerID:    "user1_ID",
					BorrowerID: "user2_ID",
					ISBN13:     "ISBN13",
					Title:      "book1_title",
				},
			},
			want: book.Book{
				ID:         "book1_id",
				OwnerID:    "user1_ID",
				BorrowerID: "user2_ID",
				ISBN13:     "ISBN13",
				Title:      "book1_title",
			},
			wantError: codes.OK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{}
			ctx := tt.args.ctx
			tt.fixture(ctx)
			err := r.SaveBook(ctx, tt.args.book)
			if diff := cmp.Diff(errors.Code(err), tt.wantError); diff != "" {
				t.Errorf(diff)
				t.Log(err)
			}
			got, err := r.LoadBook(ctx, tt.args.book.ID)
			if err != nil {
				t.Error(err)
				t.Log(err.Error())
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}

}
