package main

import (
	"github.com/NikitaDotsenko/go/database"
	"github.com/NikitaDotsenko/go/datastore"
	"github.com/NikitaDotsenko/go/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"time"
)

func Welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "asd!")
	}
}


func main() {
	db, err := datastore.NewDB()
	logFatal(err)

	db.LogMode(true)
	defer db.Close()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", Welcome())
	e.GET("/users", handler.GetUsers(db))


	if err := e.Start(":3000"); err != nil {
		log.Fatalln(err)
	}

}
func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
