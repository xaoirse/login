package handler

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

// func permissionCheck(h echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		/*
// 			Get func(r *Request, name string) (*Session, error)
// 			Get returns a session for the given name after adding it to the registry.
// 			It returns a new session if the sessions doesn't exist. Access IsNew on the session to check if it is an existing session or a new one.
// 			It returns a new session and an error if the session exists but could not be decoded.
// 			Get returns a session for the given name after adding it to the registry.
// 			See CookieStore.Get().
// 			Get registers and returns a session for the given name and session store.
// 			It returns a new session if there are no sessions registered for the name.
// 		*/
// 		sess, ok := session.Get("mySession", c)
// 		if ok != nil {
// 			return c.Redirect(http.StatusSeeOther, "/login")
// 		} else if auth, ok := sess.Values["Valid"]; auth == "yes" || !ok {
// 			fmt.Println(auth, ok)
// 			return c.Redirect(http.StatusSeeOther, "/login")
// 		}
// 		return h(c)
// 		/* or
// 		err := mid(c)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 		*/
// 	}
// }
func validationCheck(str string, format ...string) bool {
	for _, f := range format {
		switch f {
		case "required":
			if len(str) < 1 || len(str) > 255 {
				fmt.Println("Invalid length")
				return false
			}

		case "number":
			_, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Invalid number")
				return false
			}
			// check range of number

		case "english":
			if m, _ := regexp.MatchString("^[a-zA-Z]+$", str); !m {
				fmt.Println("Invalid character")
				return false
			}

		case "email":
			if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, str); !m {
				fmt.Println("Invalid email address")
				return false
			}

		default:
			panic("Invalid validationChecker format string!")
		}
	}

	return true
}

func TokenCheck(mid echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("mySession", c)
		/*
			Get func(r *Request, name string) (*Session, error)
			Get returns a session for the given name after adding it to the registry.
			It returns a new session if the sessions doesn't exist. Access IsNew on the session to check if it is an existing session or a new one.
			It returns a new session and an error if the session exists but could not be decoded.
			Get returns a session for the given name after adding it to the registry.
			See CookieStore.Get().
			Get registers and returns a session for the given name and session store.
			It returns a new session if there are no sessions registered for the name.
		*/
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		} else if auth, ok := sess.Values["foo"]; auth == "no" || !ok {
			fmt.Println("auth:", auth, "ok:", ok)
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		fmt.Println("here")
		return mid(c)
		/* or
		err := mid(c)
		if err != nil {
			return err
		}
		 return nil
		*/
	}
}

func Index(c echo.Context) error {
	sess, _ := session.Get("mySession", c)

	return c.Render(http.StatusOK, "index.html", sess.Values["foo"])
}

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

func Dashboard(c echo.Context) error {
	sess, _ := session.Get("mySession", c)

	return c.Render(http.StatusOK, "dashboard.html", sess.Values["name"])
}

func Login(c echo.Context) error {
	sess, _ := session.Get("mySession", c)

	username := c.FormValue("username")
	password := c.FormValue("password")
	println(sess.Values["formToken"])
	println(c.FormValue("validity"))
	// TODO insert a func for validity check for log
	if sess.Values["formToken"] == c.FormValue("validity") &&
		validationCheck(username, "required") &&
		validationCheck(password, "required") {
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

// TODO logs most be printed in fmt.errorf
func Upload(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return err
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer src.Close()

	// Destination
	// TODO check size and format and validation
	// TODO manage name of files
	// TODO defer close fix
	dst, err := os.Create(file.Filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	// TODO redirect to uploaded page or move this handler
	return c.Redirect(http.StatusSeeOther, "/dashboard")
}

func Logout(c echo.Context) error {
	sess, _ := session.Get("mySession", c)
	sess.Values["foo"] = "no"
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/")
}
