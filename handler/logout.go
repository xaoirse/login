package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	sess, _ := session.Get("mySession", c)
	sess.Values["foo"] = "no"
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/")
}
