package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xaoirse/logbook/controller"
	"github.com/xaoirse/logbook/graph"
	"github.com/xaoirse/logbook/graph/generated"
	mymidd "github.com/xaoirse/logbook/middleware"
)

// New return a new *Echo
func New(db *gorm.DB, secret *string) *echo.Echo {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	e := echo.New()
	// TODO random secret generator
	// Note: Don't store your key in your source code. Pass it via an
	// environmental variable, or flag (or both), and don't accidentally commit it
	// alongside your code. Ensure your key is sufficiently random - i.e. use Go's
	// crypto/rand or securecookie.GenerateRandomKey(32) and persist the result.
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(*secret))))
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Secure())
	// TODO uncomment for release or write better one
	// e.Use(middleware.Logger())

	// TODO a middleware for flood check
	// Root
	e.GET("/home/", controller.Index(db))
	e.GET("/login/", controller.LoginPage(db))
	e.POST("/login/", controller.Login(db))
	e.POST("/logout/", controller.Logout(db))

	// gqlgen
	e.GET("/", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))
	e.POST("/query/", echo.WrapHandler(srv))

	// Dashboard
	dash := e.Group("/dashboard")
	dash.Use(mymidd.SessionChecker)
	dash.GET("/", controller.Dashboard(db))
	dash.POST("/upload/", controller.Upload(db))

	return e
}
