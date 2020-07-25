package handler

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	sess, _ := session.Get("mySession", c)

	username := c.FormValue("username")
	password := c.FormValue("password")
	println(sess.Values["formToken"])
	println(c.FormValue("validity"))
	// TODO insert a func for validity check for log
	if sess.Values["formToken"] == c.FormValue("validity") &&
		ValidationCheck(username, "required") &&
		ValidationCheck(password, "required") {
		if username == "s" && password == "s" {

			sess.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   600,
				HttpOnly: true,
			}
			sess.Values["foo"] = "bar"
			sess.Values["name"] = "<script>alert('hello')</script>"
			sess.Save(c.Request(), c.Response())

			return c.Redirect(http.StatusSeeOther, "/dashboard/")
		}
	}

	return c.Redirect(http.StatusSeeOther, "/login/")
}
