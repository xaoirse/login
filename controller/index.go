package controller

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	sess, _ := session.Get("mySession", c)

	return c.Render(http.StatusOK, "index.html", sess.Values["foo"])
}
