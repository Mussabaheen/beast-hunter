// Package battle implements the core battle logic between the chracters
package battle

import (
	"errors"

	"github.com/sirupsen/logrus"
)

var (
	// ErrNoWinnerFound represents the case in which no character won
	ErrNoWinnerFound = errors.New("no winner, both characters still stand the ground")
)

// Battle is the entry point for the battle
func Battle(hero, beast *Character) (*Character, error) {
	logrus.Info("*******************************")
	logrus.Info(hero)
	logrus.Info(beast)
	logrus.Info("*******************************")
	logrus.Info("Let the battle Begin!!")
	attacker, defender := firstAttacker(hero, beast)
	for i := 1; i <= 20; i++ {
		attack(attacker, defender)
		if hero.Health <= 0 {
			return beast, nil
		}
		if beast.Health <= 0 {
			return hero, nil
		}
		attacker, defender = defender, attacker
	}
	return nil, ErrNoWinnerFound
}

func firstAttacker(hero, beast *Character) (*Character, *Character) {
	if hero.Speed > beast.Speed {
		return hero, beast
	} else if hero.Speed < beast.Speed {
		return beast, hero
	} else if hero.Luck < beast.Luck {
		return beast, hero
	} else {
		return hero, beast
	}
}

func attack(attacker, defender *Character) {
	strikeCount := 1
	for _, skill := range attacker.Skills {
		if skill.ShouldApply() {
			if skill.StrikeCountModifier != 0.0 {
				strikeCount = int(skill.StrikeCountModifier)
				logrus.Info(skill.Name, " Activated by ", attacker.Name)
			}
		}
	}
	for i := 0; i < strikeCount; i++ {
		if defender.IsLucky() {
			continue
		}
		damage := (attacker.Strength - defender.Defence)
		logrus.Info("Damage : ", damage)

		for _, skill := range defender.Skills {
			if skill.ShouldApply() {
				if skill.DamageModifier != 0.0 {
					damage = damage / int(skill.DamageModifier)
					logrus.Info(skill.Name+" Activated by ", defender.Name, " damage reduced to ", damage)
				}
			}
		}
		defender.Health = defender.Health - damage
		logrus.Info("Defender ", defender)
		logrus.Info("Attacker ", attacker)
	}
}
