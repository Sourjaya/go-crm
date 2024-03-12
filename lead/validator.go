package lead

import (
	"errors"
	"regexp"
)

type Validator interface {
	isValid() (bool, error)
}
type Phone struct {
	phone string
}
type Email struct {
	email string
}

func (e Email) isValid() (bool, error) {
	r := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	result := r.MatchString(e.email)
	var err error = nil
	if !result {
		err = errors.New("error validating email number")
	}
	return result, err
}
func (p Phone) isValid() (bool, error) {
	r := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	result := r.MatchString(p.phone)
	var err error = nil
	if !result {
		err = errors.New("error validating phone number")
	}
	return result, err
}
