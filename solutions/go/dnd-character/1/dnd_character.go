package dndcharacter

import(
    "math"
    "math/rand/v2"
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

func min(a, b, c, d int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	if d < m {
		m = d
	}
	return m
}


// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score)-10)/2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	n1 := rand.IntN(6) + 1
    n2 := rand.IntN(6) + 1
    n3 := rand.IntN(6) + 1
    n4 := rand.IntN(6) + 1
    return n1 + n2 + n3 + n4 - min(n1,n2,n3,n4)
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	strength := Ability()
    dexterity := Ability()
    constitution := Ability()
    intelligence := Ability()
    wisdom := Ability()
    charisma := Ability()
    hitpoints := 10 + Modifier(constitution)
    return Character{
        Strength: strength,
        Dexterity: dexterity,
        Constitution: constitution,
        Intelligence: intelligence,
        Wisdom: wisdom,
        Charisma: charisma,
        Hitpoints: hitpoints,
    }
}
