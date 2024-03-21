package users

import (
	"errors"
	"net/mail"
	"regexp"
)

func ValidatePassword(password string) (bool, error) {
	if len(password) < 8 {
		return false, errors.New("password must be at least 8 characters long")
	}
	if len(password) > 60 {
		return false, errors.New("password must be at most 60 characters long")
	}
	if match, _ := regexp.MatchString(`[0-9]`, password); !match {
		return false, errors.New("password must contain at least one number")
	}
	if match, _ := regexp.MatchString(`[a-zA-Z]`, password); !match {
		return false, errors.New("password must contain at least one letter")
	}
	return true, nil
}

func ValidateUsername(uname string) (bool, error) {
	if len(uname) < 1 || len(uname) > 30 {
		return false, errors.New("username must be between 1 and 30 characters long")
	}
	if uname[0] == '_' || uname[len(uname)-1] == '_' {
		return false, errors.New("username cannot start or end with an underscore")
	}
	if match, _ := regexp.MatchString(`^[a-zA-Z0-9_.]*$`, uname); !match {
		return false, errors.New("username can only contain alphanumeric characters (letters A-Z, numbers 0-9) and periods")
	}
	return true, nil
}

func ValidateEmail(email string) (bool, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, errors.New("email is invalid")
	}
	return true, nil
}
