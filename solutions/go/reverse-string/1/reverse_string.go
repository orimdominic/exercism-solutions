package reverse

func Reverse(input string) string {
	s := ""
	for _, c := range input {
		s = string(c) + s
	}
	return s
}
