package usecase

import (
	"github.com/google/go-cmp/cmp"
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/domain/repositories"
	"github.com/yuonoda/bookspace/app/domain/repositories/mock"
	"github.com/yuonoda/bookspace/app/errors"
	"github.com/yuonoda/bookspace/app/errors/codes"
	"github.com/yuonoda/bookspace/app/usecase/dto"
	"testing"
)

func TestUsecase_CreateUser(t *testing.T) {
	type fields struct {
		repositories func(t *testing.T) repositories.Repository
	}
	type args struct {
		dtoUser dto.User
	}
	tests := []struct {
		name      string
		args      args
		fields    fields
		wantError codes.Code
		want      dto.User
	}{
		{
			name: "pass",
			args: args{
				dto.User{
					Name: "Taro",
				},
			},
			fields: fields{
				repositories: func(t *testing.T) repositories.Repository {
					return &mock.RepositoryMock{
						SaveUserFunc: func(user user.User) error {
							return nil
						},
					}
				},
			},
			wantError: codes.OK,
			want: dto.User{
				Name: "Taro",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository: tt.fields.repositories(t),
			}
			got, err := u.CreateUser(tt.args.dtoUser)
			if diff := cmp.Diff(errors.Code(err), tt.wantError); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
