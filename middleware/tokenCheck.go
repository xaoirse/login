package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

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
