package controller

import (
	"fmt"
	"regexp"
	"strconv"
)

// FieldValidationCheck for "required", "number", "english", "email"
// and panic for else
func FieldValidationCheck(str string, format ...string) bool {
	// TODO should return error to user
	for _, f := range format {
		switch f {
		case "required":
			if len(str) < 1 || len(str) > 255 {
				fmt.Println("Invalid length")
				return false
			}
		// TODO check range of number
		case "number":
			_, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Invalid number")
				return false
			}

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
