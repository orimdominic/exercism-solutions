package acronym

import (
	"unicode"
)

// Abbreviate converts a string to its acronym
func Abbreviate(s string) string {
	useNext := true
	abbr := make([]rune, 0)

	for _, r := range s {
		if useNext && !isLetter(r) {
			continue
		}

		if useNext {
			abbr = append(abbr, unicode.ToUpper(r))
			useNext = false
		}

		if r == ' ' || r == '-' {
			useNext = true
		}
	}

	return string(abbr)
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
