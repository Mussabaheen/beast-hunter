package battle

import (
	"errors"
	"math/rand"
)

var (
	// ErrSkillValuesNotEqual represents that skill values are not equal
	ErrSkillValuesNotEqual = errors.New("skill values are not equal")
)

// Skill represents the abilities of the character
type Skill struct {
	Name                string  // name of the skill
	StrikeCountModifier float64 // strike count modifier value for the skill
	DamageModifier      float64 // damage modifier value for the skill
	Chance              float64 // chance of the skill to occur
}

// ShouldApply checks for the probability of the skill
func (s *Skill) ShouldApply() bool {
	return rand.Intn(101) <= int(s.Chance)
}

// NewCharacterSkills returns the Skill list for the character
func NewCharacterSkills(strikeCountModifier, damageModifier, chance []float64, name []string) ([]Skill, error) {
	var characterSkills []Skill
	if len(name) == len(strikeCountModifier) &&
		len(strikeCountModifier) == len(damageModifier) &&
		len(damageModifier) == len(chance) {
		for index, skillName := range name {
			characterSkills = append(characterSkills, Skill{
				Name:                skillName,
				StrikeCountModifier: strikeCountModifier[index],
				DamageModifier:      damageModifier[index],
				Chance:              chance[index],
			})
		}
		return characterSkills, nil
	}

	return characterSkills, ErrSkillValuesNotEqual
}
