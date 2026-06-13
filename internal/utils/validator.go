package utils

import (
	"errors"
	"net/mail"
	"net/url"
	"regexp"
)

func IsValidURL(rawURL string) bool {
	parsedURL, err := url.ParseRequestURI(rawURL)

	if err != nil {
		return false
	}

	return parsedURL.Scheme != "" &&
		parsedURL.Host != ""
}

func ValidateUsername(username string) error {
    if len(username) < 3 || len(username) > 30 {
        return errors.New("username must be between 3 and 30 characters")
    }

    matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, username)

    if !matched {
        return errors.New("username can only contain letters, numbers and underscores")
    }

    return nil
}

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return errors.New("invalid email format")
	}

	return nil
}


func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	return nil
}