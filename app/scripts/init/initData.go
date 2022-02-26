package main

import (
	"context"
	infra "github.com/yuonoda/bookspace/app/infrastracture/firestore"
	"github.com/yuonoda/bookspace/app/lib/ctxlib"
	"log"
	"os"
)

func main() {
	// CLI用のハンドラーを作ったほうが楽そう
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8812")
	os.Setenv("GCLOUD_PROJECT_ID", "local-project")
	ctx := context.Background()
	ctx = ctxlib.SetDB(ctx)
	store := ctxlib.GetDB(ctx)
	_, err := store.Collection(infra.UsersCollectionName).Doc("2zDk5t20kKM4wc2WCrZOe").Set(ctx, map[string]interface{}{
		"ID":   "2zDk5t20kKM4wc2WCrZOe",
		"name": "Taro",
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = store.Collection(infra.BooksCollectionName).Doc("9cw_QRIFHFq1EMPoy4JL5").Set(ctx, map[string]interface{}{
		"ID":         "9cw_QRIFHFq1EMPoy4JL5",
		"OwnerID":    "2zDk5t20kKM4wc2WCrZOe",
		"Title":      "book1_title",
		"BorrowerID": "",
		"ISBN13":     "",
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = store.Collection(infra.BooksCollectionName).Doc("rolrxsqkS2zgaEDbAAhit").Set(ctx, map[string]interface{}{
		"ID":         "rolrxsqkS2zgaEDbAAhit",
		"OwnerID":    "2zDk5t20kKM4wc2WCrZOe",
		"Title":      "book2_title",
		"BorrowerID": "",
		"ISBN13":     "",
	})
	if err != nil {
		log.Fatal(err)
	}
}
