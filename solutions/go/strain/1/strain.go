package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](list []T, predicate func(T) bool) []T {
	kept := []T{}

	for _, item := range list {
		if predicate(item) {
			kept = append(kept, item)
		}
	}

	return kept
}

func Discard[T interface{}](list []T, predicate func(T) bool) []T {
	discarded:=  []T{}

	for _, item := range list {
		if !predicate(item) {
			discarded = append(discarded, item)
		}
	}

	return discarded
}