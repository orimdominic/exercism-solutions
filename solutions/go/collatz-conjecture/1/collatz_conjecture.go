package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("expects positive number greater than 0")
	}

	if n == 1 {
		return 0, nil
	}

	var count = 0

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			count++
			continue
		}

		if n%2 == 1 {
			n = (3 * n) + 1
			count++
			continue
		}
	}

	return count, nil
}
