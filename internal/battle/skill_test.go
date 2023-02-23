package battle

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkill_ShouldApplyTrue(t *testing.T) {
	rand.Seed(0) // Fixing the seed value for tests that depend on RNG to produce deterministic results for testing

	skill := Skill{
		Name:                "Rapid Strike",
		Chance:              15,
		StrikeCountModifier: 2.0,
	}

	assert.True(t, skill.ShouldApply(), "must be true")
}

func TestSkill_ShouldApply(t *testing.T) {
	rand.Seed(0) // Fixing the seed value for tests that depend on RNG to produce deterministic results for testing

	skill := Skill{
		Name:                "Rapid Strike",
		Chance:              10,
		StrikeCountModifier: 2.0,
	}

	assert.False(t, skill.ShouldApply(), "must be false")
}

func TestSkill_NewCharacterSkills(t *testing.T) {

	skillNames := []string{"3xKamehameHA", "sprit bomb"}
	strikeCountModifier := []float64{3.0, 0.0}
	damageModifier := []float64{0.0, 3.0}
	chanceSkill := []float64{20.0, 30.0}

	skills, err := NewCharacterSkills(strikeCountModifier, damageModifier, chanceSkill, skillNames)
	assert.Nil(t, err, "error must be nil")

	assert.Equal(t, skills[0].Name, skillNames[0], "must be equal")
	assert.Equal(t, skills[1].Name, skillNames[1], "must be equal")
	assert.Equal(t, skills[0].StrikeCountModifier, strikeCountModifier[0], "must be equal")
	assert.Equal(t, skills[1].StrikeCountModifier, strikeCountModifier[1], "must be equal")
	assert.Equal(t, skills[0].DamageModifier, damageModifier[0], "must be equal")
	assert.Equal(t, skills[1].DamageModifier, damageModifier[1], "must be equal")
	assert.Equal(t, skills[0].Chance, chanceSkill[0], "must be equal")
	assert.Equal(t, skills[0].Chance, chanceSkill[0], "must be equal")
}
func TestSkill_NewCharacterSkillsReturnsErrSkillValuesNotEqual(t *testing.T) {

	skillNames := []string{"3xKamehameHA", "sprit bomb"}
	strikeCountModifier := []float64{3.0, 0.0}
	damageModifier := []float64{0.0, 3.0}
	chanceSkill := []float64{20.0}

	_, err := NewCharacterSkills(strikeCountModifier, damageModifier, chanceSkill, skillNames)
	assert.ErrorIs(t, err, ErrSkillValuesNotEqual, "error must be nil")

}
