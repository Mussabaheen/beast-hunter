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
	rand.Seed(time.Now().UnixNano())

	env := ""
	if envVar := os.Getenv("ENV"); envVar != "" {
		env = strings.ToLower(env)
	}

	err := config.LoadConfig(env, "./config")
	if err != nil {
		logrus.Panic("unable to load config")
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	heroSkills, err := battle.NewCharacterSkills(config.Global.HeroSkillConfig.StrikeCountModifier,
		config.Global.HeroSkillConfig.DamageModifier, config.Global.HeroSkillConfig.SkillChance, config.Global.HeroSkillConfig.Name)
	if err != nil {
		logrus.Panic(err)
	}

	beastSkills, err := battle.NewCharacterSkills(config.Global.BeastSkillConfig.StrikeCountModifier,
		config.Global.BeastSkillConfig.DamageModifier, config.Global.BeastSkillConfig.SkillChance, config.Global.BeastSkillConfig.Name)
	if err != nil {
		logrus.Panic(err)
	}

	heroStats := battle.NewRandomStats(config.Global.HeroConfig.HealthMin, config.Global.HeroConfig.HealthMax, config.Global.HeroConfig.StrengthMin,
		config.Global.HeroConfig.StrengthMax, config.Global.HeroConfig.DefenceMin, config.Global.HeroConfig.DefenceMax, config.Global.HeroConfig.SpeedMin, config.Global.HeroConfig.SpeedMax,
		config.Global.HeroConfig.LuckMin, config.Global.HeroConfig.LuckMax)
	beastStats := battle.NewRandomStats(config.Global.BeastConfig.HealthMin, config.Global.BeastConfig.HealthMax, config.Global.BeastConfig.StrengthMin,
		config.Global.BeastConfig.StrengthMax, config.Global.BeastConfig.DefenceMin, config.Global.BeastConfig.DefenceMax, config.Global.BeastConfig.SpeedMin, config.Global.BeastConfig.SpeedMax,
		config.Global.BeastConfig.LuckMin, config.Global.BeastConfig.LuckMax)

	hero := &battle.Character{Name: "Netulus", Stats: *heroStats, Skills: heroSkills}
	beast := &battle.Character{Name: "Beast", Stats: *beastStats, Skills: beastSkills}

	winner, err := battle.Battle(hero, beast)

	if err != nil {
		logrus.Error(err.Error())
	}

	logrus.Info(winner.Name, " has won the battle")
}
