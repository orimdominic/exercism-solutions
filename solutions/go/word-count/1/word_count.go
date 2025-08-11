package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	var mappings Frequency = make(Frequency)
	currStr := ""

	for i := 0; i < len(phrase); i++ {
		ch := phrase[i]

		if unicode.IsSpace(rune(ch)) {
			mappings = addToMap(currStr, mappings)
			currStr = ""
			continue
		}

		if string(ch) == "'" {
			currStr += string(ch)
			continue
		}

		if !unicode.IsLetter(rune(ch)) && !unicode.IsDigit(rune(ch)) {
			mappings = addToMap(currStr, mappings)
			currStr = ""
			continue
		}

		currStr += string(ch)
	}

	mappings = addToMap(currStr, mappings)

	return mappings
}

func addToMap(key string, mappings Frequency) Frequency {
	key = strings.TrimSuffix(key, "'")
	key = strings.TrimPrefix(key, "'")

	if len(key) == 0 {
		return mappings
	}

	mappings[key]++

	return mappings
}
