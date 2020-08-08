package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Dashboard is hanldler for GET /dashboard/
func Dashboard(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		sess, _ := session.Get("Session", c)
		return c.Render(http.StatusOK, "dashboard.html", sess.Values["username"])
	}
}

// Upload is handler for POST /upload/
// TODO logs most be printed in fmt.errorf
func Upload(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {

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
}
