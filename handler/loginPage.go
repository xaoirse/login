package handler

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type loginPageData struct {
	Token string
}

func LoginPage(c echo.Context) error {
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
