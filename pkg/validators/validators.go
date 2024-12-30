package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var phoneRegex = regexp.MustCompile(`^1\\d{10}$`)

func Phone(fl validator.FieldLevel) bool {
	phone, ok := fl.Field().Interface().(string)
	if ok {
		return phoneRegex.MatchString(phone)
	}
	return true
}
