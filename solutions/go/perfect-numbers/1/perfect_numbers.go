package perfect

import (
	"errors"
)

var ErrOnlyPositive = errors.New("requires numbers greater than 0")

// Define the Classification type here.
type Classification string

const ClassificationDeficient Classification = "deficient"
const ClassificationPerfect Classification = "perfect"
const ClassificationAbundant Classification = "abundant"
const Unclassified Classification = ""

func Classify(n int64) (Classification, error) {
	if n < 1 {
		return Unclassified, ErrOnlyPositive
	}

	half := n / 2
	factorHash := map[int64]bool{}

	for i := 1; i < int(half); i++ {
		if _, ok := factorHash[int64(i)]; ok {
			continue
		}

		if n%int64(i) == 0 {
			factorHash[int64(i)] = true
			factorHash[n/int64(i)] = true
		}
	}

	var sum int64
	for k := range factorHash {
		if k == n {
			continue
		}

		sum = sum + k
	}

	switch {
	case n == sum:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
