package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Index(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		sess, _ := session.Get("mySession", c)

		return c.Render(http.StatusOK, "index.html", sess.Values["foo"])
	}
}
