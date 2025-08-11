package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	normalised := ""

	for _, r := range pt {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			normalised += strings.ToLower(string(r))
		}
	}
	numOfRows, numOfCols := getRowCol(normalised)

	chunks := make([]string, numOfCols)
	for i, r := range normalised {
		pos := i % numOfCols
		chunks[pos] = chunks[pos] + string(r)
	}

	var sb strings.Builder

	for i, chunk := range chunks {
		sb.Write([]byte(chunk))

		if len(chunk) < numOfRows {
			sb.Write([]byte(strings.Repeat(" ", numOfRows-len(chunk))))
		}

		if i < len(chunks)-1 {
			sb.Write([]byte(" "))
		}

	}

	return sb.String()
}

// getRowCol returns the row and the column in a tuple (row, col)
func getRowCol(normalised string) (int, int) {
	min := int(math.Sqrt(float64(len(normalised))))
	if min*min == len(normalised) {
		return min, min
	}

	if min*(min+1) >= len(normalised) {
		return min, min + 1

	}

	return min + 1, min + 1
}
