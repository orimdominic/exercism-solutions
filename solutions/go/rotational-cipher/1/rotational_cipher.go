package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	s := ""
	for _, r := range plain {
		if !unicode.IsLetter(r) {
			s += string(r)
			continue
		}

		shiftBy := rune(shiftKey % 26)

		if unicode.IsLower(r) && r+shiftBy > 'z' {
			shiftBy = r + shiftBy - 26
			s += string(shiftBy)
			continue
		}

		if unicode.IsUpper(r) && r+shiftBy > 'Z' {
			shiftBy = r + shiftBy - 26
			s += string(shiftBy)
			continue
		}

		s += string(r + shiftBy)
	}

	return s
}
