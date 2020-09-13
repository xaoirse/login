package controller

import (
	"net/http"

	"github.com/xaoirse/logbook/graph/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type loginPageData struct {
	Validity string
}

// LoginPage is handler for GET /login/
func LoginPage(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		s := model.Session{}
		sessionToken := s.New(c, db)
		lpd := loginPageData{Validity: sessionToken}
		return c.Render(http.StatusOK, "login.html", lpd)
	}
}

// Login is handler for POST /login/
func Login(db *gorm.DB) func(echo.Context) error {

	return func(c echo.Context) error {

		// TODO should POST hashed password
		username := c.FormValue("username")
		password := c.FormValue("password")

		// TODO insert a func for validity check for log
		if model.IsSessionValid(c, true) &&

			FieldValidationCheck(username, "required") &&
			FieldValidationCheck(password, "required") {

			if username == "s" && password == "s" {
				s := model.Session{
					Username: "s",
				}
				s.New(c, db)
				return c.Redirect(http.StatusSeeOther, "/dashboard/")
			}
		}

		return c.Redirect(http.StatusSeeOther, "/login/")
	}
}
