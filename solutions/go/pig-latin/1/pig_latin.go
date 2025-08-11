package piglatin

import (
	"regexp"
	"strings"
)

func Sentence(sentence string) string {
	words := strings.Split(sentence, " ")
	var result []string

	for _, w := range words {
		if startsWithVowelOrXROrYT(w) {
			result = append(result, w+"ay")
			continue
		}

		quIndex := getQUIndex(w)
		if quIndex != -1 {
			result = append(result, w[quIndex+2:]+w[0:quIndex+2]+"ay")
			continue
		}

		firstVowelIndex := getFirstVowelIndex(w)
		if firstVowelIndex != -1 {
			result = append(result, w[firstVowelIndex:]+w[0:firstVowelIndex]+"ay")
			continue
		}

		yIndex := getInnerYIndex(w)

		if yIndex != -1 {
			result = append(result, w[yIndex:]+w[0:yIndex]+"ay")
			continue
		}

	}

	return strings.Join(result, " ")
}

func startsWithVowelOrXROrYT(s string) bool {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	if _, ok := vowels[rune(s[0])]; ok {
		return true
	}

	// if len(s) == 1 {
	// 	return true
	// }

	if s[0:2] == "xr" || s[0:2] == "yt" {
		return true
	}

	return false
}

// getQUIndex checks if string s starts with "qu" or has "qu" as s substring
func getQUIndex(s string) int {
	re := regexp.MustCompile(`qu`)
	loc := re.FindStringIndex(s)
	if loc == nil {
		return -1
	}

	return loc[0]
}

// getInnerYIndex checks if a string starts with one or more consonants followed by "y"
func getInnerYIndex(s string) int {
	for i := 1; i < len(s); i++ {
		if s[i] == 'y' {
			return i
		}
	}

	return -1
}

func getFirstVowelIndex(s string) int {
	re := regexp.MustCompile(`(a|e|i|o|u)`)
	loc := re.FindStringIndex(s)

	if loc == nil {
		return -1
	}

	return loc[0]
}
