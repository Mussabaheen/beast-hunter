package battle

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBattle_BeastWinner(t *testing.T) {
	rand.Seed(0) // Fixing the seed value for tests that depend on RNG to produce deterministic results for testing

	beast := Character{
		Name: "Batman",
		Stats: Stats{
			Health:   72,
			Strength: 82,
			Defence:  21,
			Speed:    20,
			Luck:     10,
		},
	}
	hero := Character{
		Name: "Beast",
		Stats: Stats{
			Health:   40,
			Strength: 22,
			Defence:  20,
			Speed:    31,
			Luck:     20,
		},
	}
	winner, err := Battle(&hero, &beast)
	assert.Nil(t, err, "err must be nil")
	assert.Equal(t, &beast, winner)
}
func TestBattle_HeroWinner(t *testing.T) {
	rand.Seed(0) // Fixing the seed value for tests that depend on RNG to produce deterministic results for testing

	hero := Character{
		Name: "Batman",
		Stats: Stats{
			Health:   72,
			Strength: 82,
			Defence:  21,
			Speed:    20,
			Luck:     10,
		},
	}
	beast := Character{
		Name: "Beast",
		Stats: Stats{
			Health:   40,
			Strength: 22,
			Defence:  20,
			Speed:    31,
			Luck:     20,
		},
	}
	winner, err := Battle(&hero, &beast)
	assert.Nil(t, err, "err must be nil")
	assert.Equal(t, &hero, winner)
}
func TestBattle_ReturnsErrNoWinnerFound(t *testing.T) {
	rand.Seed(0) // Fixing the seed value for tests that depend on RNG to produce deterministic results for testing

	hero := Character{
		Name: "Batman",
		Stats: Stats{
			Health:   72,
			Strength: 82,
			Defence:  21,
			Speed:    20,
			Luck:     100,
		},
	}
	beast := Character{
		Name: "Beast",
		Stats: Stats{
			Health:   40,
			Strength: 22,
			Defence:  20,
			Speed:    31,
			Luck:     100,
		},
	}
	winner, err := Battle(&hero, &beast)
	assert.Nil(t, winner, "err must be nil")
	assert.ErrorIs(t, err, ErrNoWinnerFound, "must be no winner found")
}

func TestFirstAttacker_BeastAttacker(t *testing.T) {
	hero := &Character{
		Name: "hero",
		Stats: Stats{
			Speed: 20,
			Luck:  30,
		},
		Skills: []Skill{
			{
				Name:                "Rapid Strike",
				StrikeCountModifier: 2,
				Chance:              100,
			},
			{
				Name:           "Magic Shield",
				DamageModifier: 2,
				Chance:         100,
			},
		},
	}
	beast := &Character{
		Name: "beast",
		Stats: Stats{
			Speed: 50,
			Luck:  30,
		},
	}

	attacker, defender := firstAttacker(hero, beast)

	assert.Equal(t, beast, attacker, "beast must be attacker")
	assert.Equal(t, hero, defender, "hero must be defender")

}
func TestFirstAttacker_HeroAttacker(t *testing.T) {
	hero := &Character{
		Name: "hero",
		Stats: Stats{
			Speed: 70,
			Luck:  30,
		},
	}
	beast := &Character{
		Name: "beast",
		Stats: Stats{
			Speed: 50,
			Luck:  30,
		},
	}

	attacker, defender := firstAttacker(hero, beast)

	assert.Equal(t, hero, attacker, "hero must be attacker")
	assert.Equal(t, beast, defender, "beast must be defender")
}
func TestFirstAttacker_SameSpeedBeastLucky(t *testing.T) {
	hero := &Character{
		Name: "hero",
		Stats: Stats{
			Speed: 50,
			Luck:  30,
		},
	}
	beast := &Character{
		Name: "beast",
		Stats: Stats{
			Speed: 50,
			Luck:  50,
		},
	}

	attacker, defender := firstAttacker(hero, beast)

	assert.Equal(t, beast, attacker, "beast must be attacker")
	assert.Equal(t, hero, defender, "hero must be defender")
}
func TestFirstAttacker_SameSpeedHeroLucky(t *testing.T) {
	hero := &Character{
		Name: "hero",
		Stats: Stats{
			Speed: 50,
			Luck:  80,
		},
	}
	beast := &Character{
		Name: "beast",
		Stats: Stats{
			Speed: 50,
			Luck:  50,
		},
	}

	attacker, defender := firstAttacker(hero, beast)

	assert.Equal(t, hero, attacker, "hero must be attacker")
	assert.Equal(t, beast, defender, "beast must be defender")
}

func TestAttack_SkillRapidStrike(t *testing.T) {
	hero := &Character{
		Name: "hero",
		Stats: Stats{
			Health:   80,
			Strength: 72,
			Defence:  47,
			Speed:    43,
			Luck:     22,
		},
		Skills: []Skill{
			{
				Name:                "Rapid Strike",
				StrikeCountModifier: 2,
				Chance:              100,
			},
		},
	}
	beast := &Character{
		Name: "beast",
		Stats: Stats{
			Health:   62,
			Strength: 63,
			Defence:  71,
			Speed:    50,
			Luck:     0,
		},
	}

	expectedDamage := hero.Strength - beast.Defence
	expectedDamage = expectedDamage * 2
	expectedHealth := beast.Health - expectedDamage
	attack(hero, beast)

	assert.Equal(t, expectedHealth, beast.Health, "both must be equal")
}

func TestAttack_SkillMagicShield(t *testing.T) {
	hero := &Character{
		Name: "hero",
		Stats: Stats{
			Health:   80,
			Strength: 72,
			Defence:  47,
			Speed:    43,
			Luck:     22,
		},
		Skills: []Skill{
			{
				Name:           "Magic Shield",
				DamageModifier: 2,
				Chance:         100,
			},
		},
	}
	beast := &Character{
		Name: "beast",
		Stats: Stats{
			Health:   62,
			Strength: 63,
			Defence:  71,
			Speed:    50,
			Luck:     29,
		},
	}

	expectedDamage := beast.Strength - hero.Defence
	expectedDamage = expectedDamage / 2
	expectedHealth := hero.Health - expectedDamage
	attack(beast, hero)

	assert.Equal(t, expectedHealth, hero.Health, "both must be equal")
}
