package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/yuonoda/bookspace/app/handlers/rest"
	"log"
)

func main() {
	rest.Serve()
}

func initData(ctx context.Context, client *firestore.Client) {
	_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"ID":   "V1StGXR8_Z5jdHi6B-myT",
		"name": "sample2",
	})
	if err != nil {
		log.Fatal(err)
	}
}
