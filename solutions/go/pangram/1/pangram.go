package pangram

import "strings"

func IsPangram(input string) bool {
	// have a map of all chars, lowercase
	// itr through input str
	// check if a char exists in map
	// for each occurrence, increase count by 1 and then delete char from map
	// return count >= 26 or len(map) == 0
	chars := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
		"d": true,
		"e": true,
		"f": true,
		"g": true,
		"h": true,
		"i": true,
		"j": true,
		"k": true,
		"l": true,
		"m": true,
		"n": true,
		"o": true,
		"p": true,
		"q": true,
		"r": true,
		"s": true,
		"t": true,
		"u": true,
		"v": true,
		"w": true,
		"x": true,
		"y": true,
		"z": true,
	}

	for _, char := range input {
		delete(chars, strings.ToLower(string(char)))
	}

	return len(chars) == 0
}
