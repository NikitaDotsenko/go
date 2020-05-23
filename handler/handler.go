package handler

import (
	"github.com/NikitaDotsenko/go/domain/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

func GetUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var u []*model.User

		if err := db.Find(&u).Error; err != nil {
			// error handling here
			return err
		}

		return c.JSON(http.StatusOK, u)
	}
}