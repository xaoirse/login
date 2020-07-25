package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

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
