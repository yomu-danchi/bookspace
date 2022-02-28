package firestore

import (
	"context"
	"encoding/json"
	"github.com/yuonoda/bookspace/app/domain/models/book"
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"golang.org/x/xerrors"
	"google.golang.org/api/iterator"
	"log"
)

func (r *Repository) SaveBook(ctx context.Context, book book.Book) error {
	store := ctxlib.GetDB(ctx)
	_, err := store.Collection(BooksCollectionName).Doc(book.ID.String()).Set(ctx, book)
	if err != nil {
		return xerrors.Errorf("err: %w", err)
	}
	return nil
}

func (r *Repository) LoadBook(ctx context.Context, bookID book.ID) (book.Book, error) {
	store := ctxlib.GetDB(ctx)
	dsnap, err := store.Collection(BooksCollectionName).Doc(bookID.String()).Get(ctx)
	if err != nil {
		return book.Book{}, err
	}
	loadedBook, err := r.parseToBook(dsnap.Data())
	if err != nil {
		return book.Book{}, xerrors.Errorf(":%W", err)
	}
	return loadedBook, nil
}

func (r *Repository) LoadBooksOwnedBy(ctx context.Context, ownerID user.ID) (book.Books, error) {
	store := ctxlib.GetDB(ctx)
	iter := store.Collection(BooksCollectionName).Where("OwnerID", "==", ownerID).Documents(ctx)
	var fetched []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fetched = append(fetched, doc.Data())
	}
	books, err := r.parseToBooks(fetched)
	if err != nil {
		return nil, xerrors.Errorf(":%w", err)
	}
	return books, nil
}

// ジェネリクスを使ってusersと共通化したい
func (r *Repository) parseToBook(fetched map[string]interface{}) (book.Book, error) {
	bytes, err := json.Marshal(fetched)
	if err != nil {
		return book.Book{}, xerrors.Errorf("failed to parse to json : %w", err)
	}
	var b book.Book
	if err := json.Unmarshal(bytes, &b); err != nil {
		return book.Book{}, xerrors.Errorf("failed to parse from json : %w", err)
	}
	return b, err
}

// ジェネリクスを使ってusersと共通化したい
func (r *Repository) parseToBooks(fetched []map[string]interface{}) (book.Books, error) {
	bytes, err := json.Marshal(fetched)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse to json : %w", err)
	}
	var books book.Books
	if err := json.Unmarshal(bytes, &books); err != nil {
		return nil, xerrors.Errorf("failed to parse from json : %w", err)
	}
	return books, err
}
