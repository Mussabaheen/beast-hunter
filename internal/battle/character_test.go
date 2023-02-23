package battle

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacter_IsLuckyTrue(t *testing.T) {
	rand.Seed(0) // Fixing the seed value for tests that depend on RNG to produce deterministic results for testing

	character := Character{
		Name: "Hero",
		Stats: Stats{
			Luck: 30,
		},
	}

	assert.True(t, character.IsLucky(), "both must be equal")
}

func TestCharacter_IsLuckyFalse(t *testing.T) {
	rand.Seed(0) // Fixing the seed value for tests that depend on RNG to produce deterministic results for testing

	character := Character{
		Name: "Hero",
		Stats: Stats{
			Luck: 10,
		},
	}

	assert.False(t, character.IsLucky(), "both must be equal")
}
