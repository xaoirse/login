package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Logout(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {

		sess, _ := session.Get("Session", c)
		sess.Values["token"] = ""
		sess.Values["username"] = ""
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusSeeOther, "/")
	}
}
