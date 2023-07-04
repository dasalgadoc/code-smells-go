package domain

import (
	"errors"
	"regexp"
)

type EmailAddress struct {
	value string
}

func (e *EmailAddress) Value() string {
	return e.value
}

func NewEmailAddress(value string) (*EmailAddress, error) {
	err := emailIsNotEmpty(value)
	err = emailValid(value)
	if err != nil {
		return nil, err
	}

	return &EmailAddress{value: value}, nil
}

func emailValid(email string) error {
	regularExpression := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	if match, _ := regexp.MatchString(regularExpression, email); !match {
		return errors.New("email is invalid")
	}
	return nil
}

func emailIsNotEmpty(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	return nil
}
