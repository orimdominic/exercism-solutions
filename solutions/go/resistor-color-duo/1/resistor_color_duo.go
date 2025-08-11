package resistorcolorduo

import "strconv"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	bands := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}

	resistorVals := ""

	for _, color := range colors[:2] {
		resistorVals += strconv.Itoa(bands[color])
	}

	val, err := strconv.Atoi(resistorVals)
	if err != nil {
		return 0
	}

	return val
}
