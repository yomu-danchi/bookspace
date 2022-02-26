package ctxlib

import (
	"cloud.google.com/go/firestore"
	"context"
	"golang.org/x/xerrors"
	"log"
	"os"
)

const (
	DBContextKey = "DB"
)

func GetDB(ctx context.Context) *firestore.Client {
	db := ctx.Value(DBContextKey)
	client, ok := db.(*firestore.Client)
	if !ok {
		log.Fatal(xerrors.Errorf("unexpected db client type"))
	}
	return client
}

func SetDB(ctx context.Context) context.Context {
	projectID := os.Getenv("GCLOUD_PROJECT_ID")
	store, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	return context.WithValue(ctx, DBContextKey, store)
}
