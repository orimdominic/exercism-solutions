package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	var nStr string
	for _, rChar := range phoneNumber {
		if unicode.IsDigit(rChar) {
			nStr += string(rChar)
		}
	}

	if len(nStr) == 10 &&
		isValidStartDigit(string(nStr[0])) &&
		isValidStartDigit(string(nStr[3])) {
		return nStr, nil
	}

	if len(nStr) == 11 && string(nStr[0]) == "1" &&
		isValidStartDigit(string(nStr[1])) &&
		isValidStartDigit(string(nStr[4])) {
		return nStr[1:], nil
	}

	return "", errors.New("invalid number")
}

func AreaCode(phoneNumber string) (string, error) {
	nStr, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return nStr[:3], nil
}

func Format(phoneNumber string) (string, error) {
	nStr, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", nStr[:3], nStr[3:6], nStr[6:]), nil
}

func isValidStartDigit(d string) bool {
	if d == "0" || d == "1" {
		return false
	}

	return true
}
