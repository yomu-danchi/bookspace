package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/errors"
	"github.com/yuonoda/bookspace/app/errors/codes"
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"log"
	"os"
	"testing"
)

var testStore = getTestStore()

func getTestStore() *firestore.Client {
	ctx := context.Background()
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8813")
	host := os.Getenv("FIRESTORE_EMULATOR_HOST")
	log.Printf("host: %+v", host)
	store, err := firestore.NewClient(ctx, "test-project")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("store: %+v", store)
	return store
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
				_, _, err := testStore.Collection("users").Add(ctx, map[string]interface{}{
					"ID":   "V1StGXR8_Z5jdHi6B-myT",
					"name": "sample2",
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
			r := repository{}
			ctx := tt.args.ctx
			//tt.fixtures(ctx)
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
