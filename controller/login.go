package controller

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type loginPageData struct {
	Token string
}

// LoginPage is GET handler for login page
func LoginPage(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		// TODO add random salt
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		lpd := loginPageData{Token: token}

		sess, _ := session.Get("mySession", c)
		sess.Values["formToken"] = token
		sess.Save(c.Request(), c.Response())
		return c.Render(http.StatusOK, "login.html", lpd)
	}
}

// Login is POST handler for login request
func Login(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
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
}
