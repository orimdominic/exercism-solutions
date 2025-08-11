package isbn

import (
	"unicode"
)

func IsValidISBN(isbn string) bool {
	sum := 0
	digitCount := 10

	for _, char := range isbn {
		if digitCount < 1 {
			return false
		}

		if unicode.IsPunct(char) {
			continue
		}

		if digitCount == 1 && char == 88 { // 88 == 'X'
			sum = sum + (digitCount * 10)
			digitCount--
			continue
		}

		if digitCount == 1 && !unicode.IsDigit(char) {
			return false
		}

		sum = sum + (digitCount * (int(char) - '0'))
		digitCount--
	}

	if sum == 0 || digitCount != 0 {
		return false
	}

	return sum%11 == 0
}
