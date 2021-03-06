package usecase

import (
	"context"
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
		ctx     context.Context
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
				context.Background(),
				dto.User{
					Name: "Taro",
				},
			},
			fields: fields{
				repositories: func(t *testing.T) repositories.Repository {
					return &mock.RepositoryMock{
						SaveUserFunc: func(ctx context.Context, user user.User) error {
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
			got, err := u.CreateUser(tt.args.ctx, tt.args.dtoUser)
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
func TestUsecase_GetUsers(t *testing.T) {
	type fields struct {
		repositories func(t *testing.T) repositories.Repository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantError codes.Code
		want      dto.Users
	}{
		{
			name: "pass",
			fields: fields{
				repositories: func(t *testing.T) repositories.Repository {
					return &mock.RepositoryMock{
						LoadUsersFunc: func(ctx context.Context) (user.Users, error) {
							return user.Users{
								{
									Name: "user1",
									ID:   "user1_id",
								},
								{
									Name: "user2",
									ID:   "user2_id",
								},
							}, nil
						},
					}
				},
			},
			wantError: codes.OK,
			want: dto.Users{
				{
					Name: "user1",
					ID:   "user1_id",
				},
				{
					Name: "user2",
					ID:   "user2_id",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository: tt.fields.repositories(t),
			}
			got, err := u.GetUsers(tt.args.ctx)
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
func TestUsecase_RegisterBook(t *testing.T) {
	const ownerID1 = "V1StGXR8_Z5jdHi6B-myU"
	type fields struct {
		repositories func(t *testing.T) repositories.Repository
	}
	type args struct {
		ctx     context.Context
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
				context.Background(),
				dto.Book{
					Title:   "book1",
					ISBN13:  "978-1-56619-909-4",
					OwnerID: ownerID1,
				},
			},
			fields: fields{
				repositories: func(t *testing.T) repositories.Repository {
					return &mock.RepositoryMock{
						SaveBookFunc: func(ctx context.Context, gotBook book.Book) error {
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
			got, err := u.RegisterBook(tt.args.ctx, tt.args.dtoBook)
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
func TestUsecase_BorrowBook(t *testing.T) {
	userID1 := user.ID("V1StGXR8_Z5jdHi6B-myU")
	bookID1 := book.ID("V1StGXR8_Z5jdHi6B-aaa")
	userID2 := user.ID("V1StGXR8_Z5jdHi6B-myV")
	type fields struct {
		repositories func(t *testing.T) repositories.Repository
	}
	type args struct {
		ctx        context.Context
		bookID     string
		borrowerID string
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
				bookID:     bookID1.String(),
				borrowerID: userID2.String(),
			},
			fields: fields{
				repositories: func(t *testing.T) repositories.Repository {
					return &mock.RepositoryMock{
						LoadBookFunc: func(ctx context.Context, bookID book.ID) (book.Book, error) {
							return book.Book{
								ID:      bookID1,
								Title:   "book1",
								ISBN13:  "978-1-56619-909-4",
								OwnerID: userID1,
							}, nil
						},
						LoadUserFunc: func(ctx context.Context, userID user.ID) (user.User, error) {
							return user.User{
								ID:   userID2,
								Name: "Taro",
							}, nil
						},
						SaveBookFunc: func(ctx context.Context, gotBook book.Book) error {
							wantBook := book.Book{
								ID:         bookID1,
								Title:      "book1",
								ISBN13:     book.ISBN13("978-1-56619-909-4"),
								OwnerID:    userID1,
								BorrowerID: userID2,
							}
							if diff := cmp.Diff(gotBook, wantBook); diff != "" {
								t.Fatal(diff)
							}
							return nil
						},
					}
				},
			},
			wantError: codes.OK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository: tt.fields.repositories(t),
			}
			err := u.BorrowBook(tt.args.ctx, tt.args.bookID, tt.args.borrowerID)
			if diff := cmp.Diff(errors.Code(err), tt.wantError); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
		})
	}
}
func TestUsecase_GetUser(t *testing.T) {
	userID1 := user.ID("V1StGXR8_Z5jdHi6B-myU")
	bookID1 := book.ID("V1StGXR8_Z5jdHi6B-aaa")
	bookID2 := book.ID("V1StGXR8_Z5jdHi6B-aab")
	type fields struct {
		repositories func(t *testing.T) repositories.Repository
	}
	type args struct {
		ctx    context.Context
		userID string
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
				userID: userID1.String(),
			},
			fields: fields{
				repositories: func(t *testing.T) repositories.Repository {
					return &mock.RepositoryMock{
						LoadBooksOwnedByFunc: func(ctx context.Context, userID user.ID) (book.Books, error) {
							return book.Books{
								{
									ID:         bookID1,
									OwnerID:    userID1,
									BorrowerID: "",
									ISBN13:     "book1_ISBN13",
									Title:      "book1_title",
								},
								{
									ID:         bookID2,
									OwnerID:    userID1,
									BorrowerID: "",
									ISBN13:     "book1_ISBN13",
									Title:      "book1_title",
								},
							}, nil
						},
						LoadUserFunc: func(ctx context.Context, userID user.ID) (user.User, error) {
							return user.User{
								ID:   userID1,
								Name: "user1_name",
							}, nil
						},
					}
				},
			},
			want: dto.User{
				ID:   userID1.String(),
				Name: "user1_name",
				OwnedBooks: dto.Books{
					{
						ID:         bookID1.String(),
						OwnerID:    userID1.String(),
						BorrowerID: "",
						ISBN13:     "book1_ISBN13",
						Title:      "book1_title",
					},
					{
						ID:         bookID2.String(),
						OwnerID:    userID1.String(),
						BorrowerID: "",
						ISBN13:     "book1_ISBN13",
						Title:      "book1_title",
					},
				},
			},
			wantError: codes.OK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository: tt.fields.repositories(t),
			}
			user, err := u.GetUser(tt.args.ctx, tt.args.userID)
			if diff := cmp.Diff(errors.Code(err), tt.wantError); diff != "" {
				t.Error(diff)
				t.Log(err)
			}
			if diff := cmp.Diff(user, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
