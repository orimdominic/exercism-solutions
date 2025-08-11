package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var rotations rune = 0
	var sentence = make([]string, 0)
	var word string

	for _, rChar := range s {
		if unicode.IsDigit(rChar) {
			word += string(rChar)
		}

		if unicode.IsLetter(rChar) {
			sub := 122 - unicode.ToLower(rChar) // 122 is 'z'
			mod := sub % 26
			word += string(mod + 97 + rotations)
		}

		if len(word) == 5 {
			sentence = append(sentence, word)
			word = ""
		}
	}

	sentence = append(sentence, strings.Trim(word, " "))

	return strings.Trim(strings.Join(sentence, " "), " ")
}
