package battle

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkill_ShouldApplyTrue(t *testing.T) {
	rand.Seed(0) // Fixing the seed value for tests that depend on RNG to produce deterministic results for testing

	skill := Skill{
		Name:             "Rapid Strike",
		Chance:           15,
		StrengthModifier: 2.0,
	}

	assert.True(t, skill.ShouldApply(), "must be true")
}

func TestSkill_ShouldApply(t *testing.T) {
	rand.Seed(0) // Fixing the seed value for tests that depend on RNG to produce deterministic results for testing

	skill := Skill{
		Name:             "Rapid Strike",
		Chance:           10,
		StrengthModifier: 2.0,
	}

	assert.False(t, skill.ShouldApply(), "must be false")
}
