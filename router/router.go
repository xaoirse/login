package router

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xaoirse/logbook/handler"
)

// New return a new *Echo
func New() *echo.Echo {
	e := echo.New()
	// TODO random secret generator
	// Note: Don't store your key in your source code. Pass it via an
	// environmental variable, or flag (or both), and don't accidentally commit it
	// alongside your code. Ensure your key is sufficiently random - i.e. use Go's
	// crypto/rand or securecookie.GenerateRandomKey(32) and persist the result.
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Secure())
	// e.Use(middleware.Logger())
	e.GET("/", handler.Index)
	e.GET("/login/", handler.LoginPage)
	e.POST("/login/", handler.Login)
	e.POST("/logout/", handler.Logout)

	dash := e.Group("/dashboard")
	dash.Use(handler.TokenCheck)
	dash.GET("/", handler.Dashboard)
	dash.POST("/upload/", handler.Upload)
	return e
}
