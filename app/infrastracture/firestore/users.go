package firestore

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/errors"
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"golang.org/x/xerrors"
	"google.golang.org/api/iterator"
	"log"
)

func (r *Repository) SaveUser(ctx context.Context, user user.User) error {
	store := ctxlib.GetDB(ctx)
	_, err := store.Collection(UsersCollectionName).Doc(user.ID.String()).Set(ctx, user)
	if err != nil {
		return xerrors.Errorf("err: %w", err)
	}
	return nil
}

func (r *Repository) LoadUser(ctx context.Context, userID user.ID) (user.User, error) {
	store := ctxlib.GetDB(ctx)
	iter := store.Collection(UsersCollectionName).Where("ID", "==", userID).Documents(ctx)
	var fetched []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return user.User{}, xerrors.Errorf(":%w", err)
		}
		fetched = append(fetched, doc.Data())
	}

	users, err := r.parseToUsers(fetched)
	if err != nil {
		return user.User{}, xerrors.Errorf(":%w", err)
	}
	if len(users) == 0 {
		return user.User{}, errors.NotFound(xerrors.Errorf("user not found, ID: %s", userID))
	}
	if len(users) >= 2 {
		return user.User{}, errors.Invalid(xerrors.Errorf("duplicated user, ID: %s", userID))
	}
	return users[0], nil
}

func (r *Repository) LoadUsers(ctx context.Context) (user.Users, error) {
	store := ctxlib.GetDB(ctx)
	iter := store.Collection("users").Documents(ctx)
	var fetched []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
		fetched = append(fetched, doc.Data())
	}

	users, err := r.parseToUsers(fetched)
	if err != nil {
		return nil, xerrors.Errorf(":%w", err)
	}
	return users, nil
}

// booksにも使えるようにしたい
func (r *Repository) parseToUsers(fetched []map[string]interface{}) (user.Users, error) {
	bytes, err := json.Marshal(fetched)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse to json : %w", err)
	}
	var users user.Users
	if err := json.Unmarshal(bytes, &users); err != nil {
		return nil, xerrors.Errorf("failed to parse from json : %w", err)
	}
	return users, err
}
