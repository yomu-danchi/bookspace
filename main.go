package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", welcome())

	err := e.Start(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
func welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to BookSpace API")
	}
}

//
//func main() {
//	dsn := "localuser:localpass@tcp(127.0.0.1:3306)/localdb?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	// Migrate the schema
//	db.Migrator().DropTable(&User{}, &Book{})
//	db.AutoMigrate(&User{}, &Book{})
//
//	user := &User{
//		Name: "小野田　祐",
//	}
//	db.Create(user)
//	fmt.Printf("user:%v\n", user)
//
//	b := &Book{
//		OwnerId: user.UserId,
//		Title:   "サンプルの本1",
//	}
//	db.Create(&b)
//}
