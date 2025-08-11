package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10) / float64(2)))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	var ability int
	lowest := 7

	for i := 0; i < 4; i++ {
		score := rand.Intn(6) + 1
		if score < lowest {
			lowest = score
		}
		ability += score
	}

	return ability - lowest
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Charisma:     Ability(),
		Wisdom:       Ability(),
	}
	c.Hitpoints = 10 + Modifier(c.Constitution)

	return c
}
