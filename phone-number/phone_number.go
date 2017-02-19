package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

const testVersion = 1

// Clean up user-entered phone numbers so that they can be sent SMS messages.
// The rules are as follows:
// - If the phone number is less than 10 digits assume that it is bad
//   number
// - If the phone number is 10 digits assume that it is good
// - If the phone number is 11 digits and the first number is 1, trim the 1
//   and use the last 10 digits
// - If the phone number is 11 digits and the first number is not 1, then
//   it is a bad number
// - If the phone number is more than 11 digits assume that it is a bad
//   number

// Given a string, return the phone number
func Number(phone string) (string, error) {
	str := strings.Map(func(r rune) rune {
		if '0' <= r && r <= '9' {
			return r
		} else {
			return -1
		}
	}, phone)

	strlen := len(str)
	if strlen == 10 {
		return str, nil
	} else if strlen == 11 {
		if str[0] == '1' {
			return str[1:], nil
		} else {
			return "", errors.New("bad number")
		}
	}
	return "", errors.New("bad number")

}

// Given a string, return the area code
// area code is the first 3 digits of a valid phone number
func AreaCode(phone string) (string, error) {
	s, err := Number(phone)
	if err == nil {
		return s[:3], nil
	} else {
		return "", err
	}
}

// Given a string, return the formated phone number if valid
// format is defined as "(xxx) xxx-xxxx"
func Format(phone string) (string, error) {
	s, err := Number(phone)
	if err == nil {
		return fmt.Sprintf("(%s) %s-%s", s[:3], s[3:6], s[6:]), nil
	} else {
		return "", err
	}
}
