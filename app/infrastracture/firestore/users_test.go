package firestore

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/errors"
	"github.com/yuonoda/bookspace/app/errors/codes"
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"log"
	"testing"
)

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
				collection := testStore.Collection(UsersCollectionName)
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
				collection := testStore.Collection(UsersCollectionName)
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
func TestRepository_SaveUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user user.User
	}
	tests := []struct {
		name      string
		args      args
		want      user.User
		wantError codes.Code
	}{
		{
			name: "success",
			args: args{
				ctx: context.WithValue(context.Background(), ctxlib.DBContextKey, testStore),
				user: user.User{
					ID:   "user1_id",
					Name: "user1_name",
				},
			},
			want: user.User{
				ID:   "user1_id",
				Name: "user1_name",
			},
			wantError: codes.OK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{}
			ctx := tt.args.ctx
			err := r.SaveUser(ctx, tt.args.user)
			if diff := cmp.Diff(errors.Code(err), tt.wantError); diff != "" {
				t.Errorf(diff)
				t.Log(err)
			}
			got, err := r.LoadUser(ctx, tt.args.user.ID)
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
