package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"os"

	//"github.com/99designs/gqlgen/graphql/handler"
	//"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/yuonoda/bookspace/app/graph"
	//"github.com/yuonoda/bookspace/app/graph/generated"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	store, err := firestore.NewClient(ctx, os.Getenv("GCLOUD_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}
	initData(ctx, store)

	//// Use the application default credentials
	//ctx := context.Background()
	//conf := &firebase.Config{ProjectID: "test-project-id"}
	//app, err := firebase.NewApp(ctx, conf)
	//if err != nil {
	//    log.Fatalln(err)
	//}
	//client, err := app.Firestore(ctx)
	//if err != nil {
	//    log.Fatalln(err)
	//}
	//defer client.Close()
	//
	//_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
	//    "first": "Ada",
	//    "last":  "Lovelace",
	//    "born":  1815,
	//})
	//if err != nil {
	//    log.Fatalf("Failed adding alovelace: %v", err)
	//}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", welcome())
	//
	//graphqlHandler := handler.NewDefaultServer(
	//	generated.NewExecutableSchema(
	//		generated.Config{Resolvers: &graph.Resolver{DB: db}}))
	//playgroundHandler := playground.Handler("GraphQl", "/query")
	//
	//e.POST("/query", func(c echo.Context) error {
	//	graphqlHandler.ServeHTTP(c.Response(), c.Request())
	//	return nil
	//})

	//e.GET("/playground", func(c echo.Context) error {
	//	playgroundHandler.ServeHTTP(c.Response(), c.Request())
	//	return nil
	//})

	if err := e.Start(":8000"); err != nil {
		log.Fatalln(err)
	}
}
func welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		store, err := firestore.NewClient(ctx, "test-project-id")
		if err != nil {
			log.Fatal(err)
		}
		iter := store.Collection("users").Documents(ctx)

		var users []map[string]interface{}
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to iterate: %v", err)
			}
			fmt.Println(doc.Data())
			users = append(users, doc.Data())
		}
		return c.JSON(http.StatusOK, users)
	}
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
