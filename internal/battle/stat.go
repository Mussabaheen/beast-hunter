package battle

import "math/rand"

// Stats represents the properties of character
type Stats struct {
	Health   int // health of the character
	Strength int // strength of the character
	Defence  int // defence of the character
	Speed    int // speed of the character
	Luck     int // luck of the character
}

// NewRandomStats takes ranges for stats and returns stats
func NewRandomStats(healthMin, healthMax, strengthMin, strengthMax,
	defenceMin, defenceMax, speedMin, speedMax, luckmin, luckMax int) *Stats {
	health := rand.Intn(healthMax-healthMin) + healthMin
	strength := rand.Intn(strengthMax-strengthMin) + strengthMin
	defense := rand.Intn(defenceMax-defenceMin) + defenceMin
	speed := rand.Intn(speedMax-speedMin) + speedMin
	luck := rand.Intn(luckMax-luckmin) + luckmin

	return &Stats{
		Health:   health,
		Strength: strength,
		Defence:  defense,
		Speed:    speed,
		Luck:     luck,
	}
}
