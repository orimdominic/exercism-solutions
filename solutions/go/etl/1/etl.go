package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out:=make(map[string]int)

	for points, alphabetSlice := range in {
		for _, alphabet:= range alphabetSlice {
			out[ strings.ToLower(alphabet)] = points
		}
	}

	return out
}
