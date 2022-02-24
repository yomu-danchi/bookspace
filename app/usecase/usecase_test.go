package usecase

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/yuonoda/bookspace/app/domain/models/book"
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
			opts := cmp.Options{
				cmpopts.IgnoreFields(dto.User{}, "ID"),
			}
			if diff := cmp.Diff(got, tt.want, opts); diff != "" {
				t.Error(diff)
			}
		})
	}
}
func TestUsecase_RegisterBook(t *testing.T) {
	const ownerID1 = "V1StGXR8_Z5jdHi6B-myU"
	type fields struct {
		repositories func(t *testing.T) repositories.Repository
	}
	type args struct {
		dtoBook dto.Book
	}
	tests := []struct {
		name      string
		args      args
		fields    fields
		wantError codes.Code
		want      dto.Book
	}{
		{
			name: "pass",
			args: args{
				dto.Book{
					Title:   "book1",
					ISBN13:  "978-1-56619-909-4",
					OwnerID: ownerID1,
				},
			},
			fields: fields{
				repositories: func(t *testing.T) repositories.Repository {
					return &mock.RepositoryMock{
						RegisterBookFunc: func(gotBook book.Book) error {
							wantBook := book.Book{
								ID:      book.ID("V1StGXR8_Z5jdHi6B-myT"),
								Title:   "book1",
								ISBN13:  book.ISBN13("978-1-56619-909-4"),
								OwnerID: ownerID1,
							}
							opts := cmp.Options{
								cmpopts.IgnoreFields(book.Book{}, "ID"),
							}
							if diff := cmp.Diff(gotBook, wantBook, opts); diff != "" {
								t.Fatal(diff)
							}

							return nil
						},
					}
				},
			},
			wantError: codes.OK,
			want: dto.Book{
				ID:      "V1StGXR8_Z5jdHi6B-myT",
				Title:   "book1",
				ISBN13:  "978-1-56619-909-4",
				OwnerID: ownerID1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository: tt.fields.repositories(t),
			}
			got, err := u.RegisterBook(tt.args.dtoBook)
			if diff := cmp.Diff(errors.Code(err), tt.wantError); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(dto.Book{}, "ID"),
			}
			if diff := cmp.Diff(got, tt.want, opts); diff != "" {
				t.Error(diff)
			}
		})
	}
}
