package middleware

import (
	"net/http"

	"github.com/xaoirse/logbook/graph/model"

	"github.com/labstack/echo/v4"
)

// SessionChecker check session validity
func SessionChecker(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if model.IsSessionValid(c, false) {
			return handler(c)
		}
		return c.Redirect(http.StatusSeeOther, "/login/")
	}
}
