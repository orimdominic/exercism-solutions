package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, count := range birdsPerDay {
		total += count
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// 1 0 -> 8
	// 2 8 -> 15
	total := 0
	start := (7 * week) - 7
	end := start + 7
	for _, count := range birdsPerDay[start:end] {
		total += count
	}

	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, count := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = count + 1
		}
	}

	return birdsPerDay
}
