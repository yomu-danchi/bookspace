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
				collection := testStore.Collection(booksCollectionName)
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
