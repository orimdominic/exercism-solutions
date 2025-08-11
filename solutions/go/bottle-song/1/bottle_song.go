package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	lines := make([]string, 0)
	for i := 0; i < takeDown; i++ {
		if len(lines) > 1 {
			lines = append(lines, "")
		}
		lines = append(lines, generateLines(startBottles-i)...)
	}

	return lines
}

func generateLines(n int) []string {
	numStrTable := map[int]string{
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
	}
	var lastLineCountStr string

	if n >= 3 {
		lastLineCountStr = strings.ToLower(numStrTable[n-1]) + " green bottles"
		return []string{
			fmt.Sprintf("%s green bottles hanging on the wall,", numStrTable[n]),
			fmt.Sprintf("%s green bottles hanging on the wall,", numStrTable[n]),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s hanging on the wall.", lastLineCountStr),
		}
	}

	if n == 2 {
		lastLineCountStr = "one green bottle"
		return []string{
			fmt.Sprintf("%s green bottles hanging on the wall,", numStrTable[n]),
			fmt.Sprintf("%s green bottles hanging on the wall,", numStrTable[n]),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s hanging on the wall.", lastLineCountStr),
		}
	}

	lastLineCountStr = "no green bottles"
	return []string{
		fmt.Sprintf("%s green bottle hanging on the wall,", numStrTable[n]),
		fmt.Sprintf("%s green bottle hanging on the wall,", numStrTable[n]),
		"And if one green bottle should accidentally fall,",
		fmt.Sprintf("There'll be %s hanging on the wall.", lastLineCountStr),
	}
}
