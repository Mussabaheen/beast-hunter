// Package main represents the entry point of the game
package main

import (
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/mussabaheen/beast-hunter/internal/battle"
	"github.com/mussabaheen/beast-hunter/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // updating the seed value everytime to get the different results for RNG

	env := ""
	if envVar := os.Getenv("ENV"); envVar != "" {
		env = strings.ToLower(env)
	}

	err := config.LoadConfig(env, "./config") // Loading configurations from config folder
	if err != nil {
		logrus.Panic("unable to load config")
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	heroStats := battle.NewRandomStats(config.Global.HeroConfig.HealthMin, config.Global.HeroConfig.HealthMax, config.Global.HeroConfig.StrengthMin,
		config.Global.HeroConfig.StrengthMax, config.Global.HeroConfig.DefenceMin, config.Global.HeroConfig.DefenceMax, config.Global.HeroConfig.SpeedMin, config.Global.HeroConfig.SpeedMax,
		config.Global.HeroConfig.LuckMin, config.Global.HeroConfig.LuckMax) // Initializing a new stat object in the given range

	beastStats := battle.NewRandomStats(config.Global.BeastConfig.HealthMin, config.Global.BeastConfig.HealthMax, config.Global.BeastConfig.StrengthMin,
		config.Global.BeastConfig.StrengthMax, config.Global.BeastConfig.DefenceMin, config.Global.BeastConfig.DefenceMax, config.Global.BeastConfig.SpeedMin, config.Global.BeastConfig.SpeedMax,
		config.Global.BeastConfig.LuckMin, config.Global.BeastConfig.LuckMax) // Initializing a new stat object in the given range

	hero := &battle.Character{Name: "Netulus", Stats: *heroStats, Skills: []battle.Skill{{
		Name:             "Rapid Strike",
		StrengthModifier: config.Global.SkillConfig.StrengthModifier,
		Chance:           config.Global.SkillConfig.StrengthChance,
	}, {
		Name:           "Magic Shield",
		DamageModifier: config.Global.SkillConfig.DamageModifier,
		Chance:         config.Global.SkillConfig.DamageChance,
	}}}

	beast := &battle.Character{Name: "Beast", Stats: *beastStats, Skills: []battle.Skill{}}

	winner, err := battle.Battle(hero, beast)

	if err != nil {
		logrus.Error(err.Error())
	}

	logrus.Info(winner.Name, " has won the battle")
}
