package main

import (
	"github.com/NikitaDotsenko/go/database"
	"github.com/NikitaDotsenko/go/datastore"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"time"
)
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Age       string    `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (User) TableName() string { return "users" }
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
