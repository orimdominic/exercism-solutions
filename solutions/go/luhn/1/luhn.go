package luhn

import (
	"strconv"
)

func Valid(id string) bool {

	digits := make([]int, 0)

	for _, c := range id {
		s := string(c)
		if s == " " {
			continue
		}

		d, err := strconv.Atoi(s)
		if err != nil {
			return false
		}

		digits = append(digits, d)
	}

	if len(digits) < 2 {
		return false
	}

	var sum = 0
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i]

		if (len(digits)-i)%2 == 0 {
			d = d * 2
		}

		if digits[i] != d && d > 9 {
			d -= 9
		}

		digits[i] = d
		sum += d
	}

	return sum%10 == 0
}
