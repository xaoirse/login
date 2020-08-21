package model

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Session struct {
	gorm.Model
	Username   string
	Role       string
	IP         string
	Token      string
	Exp        time.Time
	TryToLogin int
}

func newToken() string {
	// Generating a random token
	// TODO print as help
	salt := os.Args[1]
	crutime := time.Now().Unix()
	str := strconv.FormatInt(crutime, 10) + salt
	hash := md5.New()
	hash.Write([]byte(str))
	token := fmt.Sprintf("%x", hash.Sum(nil))
	return token
}

// New create a new session and return token
func (s *Session) New(c echo.Context) string {

	// Saving session in db
	s.Token = newToken()
	s.IP = c.RealIP()
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalln("Error in opening db:", err)
	}
	db.Create(s)

	// Saving session in response
	sess, _ := session.Get("Session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 8,
		HttpOnly: true,
	}
	sess.Values["token"] = s.Token
	sess.Values["username"] = s.Username
	sess.Values["ip"] = s.IP
	sess.Save(c.Request(), c.Response())

	return s.Token
}

// IsSessionValid returns true if session is valid
func IsSessionValid(c echo.Context, tokenCheck bool) bool {
	sess, err := session.Get("Session", c)
	/*
		Fucking diffrent between println and fmt.Println
		and type assertion in map[interface{}]interface{}
		un := sess.Values["username"].(string)
		// TODO study type assertions
	*/
	token, _ := sess.Values["token"]
	fmt.Println(sess.Values["username"].(string), token)
	if err != nil ||
		// !ok ||
		token == "" ||
		(tokenCheck && token != c.FormValue("token")) {
		// TODO check db
		return false
	}
	return true
}
