package main

import (
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
	//// TODO 環境変数にする
	//dsn := "localuser:localpass@tcp(db:3306)/localdb?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
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
		return c.String(http.StatusOK, "Welcome to BookSpace API!")
	}
}
