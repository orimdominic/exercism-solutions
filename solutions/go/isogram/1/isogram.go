package isogram

import "strings"

func IsIsogram(word string) bool {
	set := map[string]bool{}

	for _, r := range word {
		s := string(r)
		if hasChar(s, set) {
			return false
		}
		set[strings.ToLower(s)] = true
	}

	return true
}

func hasChar(s string, m map[string]bool) bool {
	if s == " " || s == "-" {
		return false
	}

	_, ok := m[strings.ToLower(s)]

	return ok
}
