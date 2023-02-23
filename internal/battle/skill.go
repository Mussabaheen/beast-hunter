package battle

import "math/rand"

// Skill represents the abilities of the character
type Skill struct {
	Name             string  // name of the skill
	StrengthModifier float64 // strength modifier value for the skill
	DamageModifier   float64 // damage modifier value for the skill
	Chance           float64 // chance of the skill to occur
}

// ShouldApply checks for the probability of the skill
func (s *Skill) ShouldApply() bool {
	return rand.Intn(101) <= int(s.Chance)
}
