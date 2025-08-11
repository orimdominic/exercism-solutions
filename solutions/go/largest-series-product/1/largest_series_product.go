package lsproduct

import (
	"errors"
	"math"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}

	if span == 0 || len(digits) == 0 {
		return 0, errors.New("span must be smaller than string length")
	}

	if span < 0 {
		return 0, errors.New("span must not be negative")
	}

	nums := make([]int, len(digits))

	for i, r := range digits {
		d, err := strconv.Atoi(string(r))

		if err != nil {
			return 0, errors.New("digits input must only contain digits")
		}

		nums[i] = d
	}

	max := int64(math.MinInt)
// 63915
	for i := 0; i < len(nums)-span+1; i++ {

		slc := nums[i : span+i]
		prod := int64(1)
		for _, n := range slc {
			prod *= int64(n)
		}

		if prod > max {
			max = prod
		}
	}

	return max, nil
}
