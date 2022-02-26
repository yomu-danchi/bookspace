package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/errors"
	"github.com/yuonoda/bookspace/app/errors/codes"
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"google.golang.org/api/iterator"
	"log"
	"os"
	"testing"
)

var testStore = getTestStore()

func getTestStore() *firestore.Client {
	ctx := context.Background()
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8813")
	store, err := firestore.NewClient(ctx, "test-project")
	if err != nil {
		log.Fatal(err)
	}
	return store
}

func deleteCollection(ctx context.Context, client *firestore.Client,
	ref *firestore.CollectionRef, batchSize int) error {

	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}

func TestRepository_LoadUsers(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name      string
		fixtures  func(ctx context.Context)
		args      args
		want      user.Users
		wantError codes.Code
	}{
		{
			name: "success",
			fixtures: func(ctx context.Context) {
				collection := testStore.Collection("users")
				deleteCollection(ctx, testStore, collection, 100)
				_, _, err := collection.Add(ctx, map[string]interface{}{
					"ID":   "V1StGXR8_Z5jdHi6B-myT",
					"name": "test_user1",
				})
				if err != nil {
					log.Fatal(err)
				}
			},
			args: args{
				ctx: context.WithValue(context.Background(), ctxlib.DBContextKey, testStore),
			},
			want: user.Users{
				{
					ID:   "V1StGXR8_Z5jdHi6B-myT",
					Name: "test_user1",
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
			got, err := r.LoadUsers(ctx)
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
func TestRepository_LoadUser(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID user.ID
	}
	tests := []struct {
		name      string
		fixtures  func(ctx context.Context)
		args      args
		want      user.User
		wantError codes.Code
	}{
		{
			name: "success",
			fixtures: func(ctx context.Context) {
				collection := testStore.Collection("users")
				deleteCollection(ctx, testStore, collection, 100)
				_, _, err := collection.Add(ctx, map[string]interface{}{
					"ID":   "V1StGXR8_Z5jdHi6B-myT",
					"name": "test_user1",
				})
				if err != nil {
					log.Fatal(err)
				}
			},
			args: args{
				ctx:    context.WithValue(context.Background(), ctxlib.DBContextKey, testStore),
				userID: user.ID("V1StGXR8_Z5jdHi6B-myT"),
			},
			want: user.User{
				ID:   "V1StGXR8_Z5jdHi6B-myT",
				Name: "test_user1",
			},
			wantError: codes.OK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{}
			ctx := tt.args.ctx
			tt.fixtures(ctx)
			got, err := r.LoadUser(ctx, tt.args.userID)
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
