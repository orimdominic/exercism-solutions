// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey returns what Bob would say in certain situations
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(strings.TrimSpace(remark)) == 0 {
		return "Fine. Be that way!"
	}

	if hasNoLetter(remark[0:len(remark)-1]) && string(remark[len(remark)-1]) == "?" {
		return "Sure."
	}

	if hasNoUpperCase(remark[0:len(remark)-1]) && string(remark[len(remark)-1]) == "?" {
		return "Calm down, I know what I'm doing!"
	}

	if hasNoUpperCase(remark[0 : len(remark)-1]) && !hasNoLetter(remark) {
		return "Whoa, chill out!"
	}

	if string(remark[len(remark)-1]) == "?" {
		return "Sure."
	}

	return "Whatever."
}

func hasNoLetter(s string) bool {
	for _, c :=range s {
		if(unicode.IsLetter(c)) {
			return false
		}
	}

	return true
}

func hasNoUpperCase(s string) bool {
	for _, c :=range s {
		if(unicode.IsLetter(c) && unicode.IsLower(c)) {
			return false
		}
	}

	return true
}
