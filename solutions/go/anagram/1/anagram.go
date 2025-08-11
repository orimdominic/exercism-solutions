package anagram

import (
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	/*
		create an empty string to hold valid candidates
		iterate over candidates. for each
		check if same length with subject
			if not same length skip
			if same length
				create hashmap of subject
				iterate over each char of the word
				if present in table, deduct and continue
				if not present in table, skip and move to next
	*/
	result := make([]string, 0)
	sameLenCandidates := make([]string, 0)
	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			continue
		}

		if len(candidate) == len(subject) {
			sameLenCandidates = append(sameLenCandidates, candidate)
		}
	}

	if len(sameLenCandidates) == 0 {
		return sameLenCandidates
	}

	table := createSubjectCharMap(strings.ToLower(subject))

	for _, candidate := range sameLenCandidates {
		candidateTable := createSubjectCharMap(strings.ToLower(candidate))
		if reflect.DeepEqual(table, candidateTable) {
			result = append(result, candidate)
		}
	}

	return result
}

func createSubjectCharMap(subject string) map[rune]int {
	table := map[rune]int{}

	for _, runeCh := range subject {
		_, ok := table[runeCh]
		if !ok {
			table[runeCh] = 1
			continue
		}
		table[runeCh]++
	}

	return table
}
