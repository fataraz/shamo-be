package helper

import (
	"regexp"
	"shamo-be/internal/shared/constant"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

// ValidateEmail ...
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// use first character is 62, 10 - 12 digit
var validNumber = regexp.MustCompile(`^((62))\d{10,13}$`)

// use first character is 08, 10 - 12 digit
var numberWith08 = regexp.MustCompile(`^((08))\d{8,11}$`)

// use first character is +62, 10 - 12 digit
var numberPlus62 = regexp.MustCompile(`^((\+62))\d{10,13}$`)

// ValidatePhoneNumber ...
func ValidatePhoneNumber(phoneNumber string) (string, error) {
	if validNumber.MatchString(phoneNumber) {
		return phoneNumber, nil
	} else if numberWith08.MatchString(phoneNumber) {
		phoneNumber = "62" + phoneNumber[1:]
		return phoneNumber, nil
	} else if numberPlus62.MatchString(phoneNumber) {
		phoneNumber = phoneNumber[1:]
		return phoneNumber, nil
	} else {
		return phoneNumber, constant.ErrorInvalidPhoneNumber
	}
}
