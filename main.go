package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"github.com/NikitaDotsenko/go/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.Welcome())

	if err := e.Start(":3000"); err != nil {
		log.Fatalln(err)
	}

}