package handler

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Dashboard(c echo.Context) error {
	sess, _ := session.Get("mySession", c)

	return c.Render(http.StatusOK, "dashboard.html", sess.Values["name"])
}
