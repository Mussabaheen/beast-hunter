// Package config handles the configuration functionality of the application
package config

import (
	"strings"

	"github.com/spf13/viper"
)

// HeroConfig have configurations for the hero stats
type HeroConfig struct {
	HealthMin   int `mapstructure:"HERO_HEALTH_MIN"`   // minimum hero health
	HealthMax   int `mapstructure:"HERO_HEALTH_MAX"`   // maximum hero health
	StrengthMin int `mapstructure:"HERO_STRENGTH_MIN"` // minimum hero strength
	StrengthMax int `mapstructure:"HERO_STRENGTH_MAX"` // maximum hero strength
	DefenceMin  int `mapstructure:"HERO_DEFENCE_MIN"`  // minimum hero defence
	DefenceMax  int `mapstructure:"HERO_DEFENCE_MAX"`  // maximum hero defence
	SpeedMin    int `mapstructure:"HERO_SPEED_MIN"`    // minimum hero speed
	SpeedMax    int `mapstructure:"HERO_SPEED_MAX"`    // maximum hero speed
	LuckMin     int `mapstructure:"HERO_LUCK_MIN"`     // minimum hero luck
	LuckMax     int `mapstructure:"HERO_LUCK_MAX"`     // maximum hero luck
}

// BeastConfig have configurations for the beast stats
type BeastConfig struct {
	HealthMin   int `mapstructure:"BEAST_HEALTH_MIN"`   // minimum Beast health
	HealthMax   int `mapstructure:"BEAST_HEALTH_MAX"`   // maximum Beast health
	StrengthMin int `mapstructure:"BEAST_STRENGTH_MIN"` // minimum Beast strength
	StrengthMax int `mapstructure:"BEAST_STRENGTH_MAX"` // maximum Beast strength
	DefenceMin  int `mapstructure:"BEAST_DEFENCE_MIN"`  // minimum Beast defence
	DefenceMax  int `mapstructure:"BEAST_DEFENCE_MAX"`  // maximum Beast defence
	SpeedMin    int `mapstructure:"BEAST_SPEED_MIN"`    // minimum Beast speed
	SpeedMax    int `mapstructure:"BEAST_SPEED_MAX"`    // maximum Beast speed
	LuckMin     int `mapstructure:"BEAST_LUCK_MIN"`     // minimum Beast luck
	LuckMax     int `mapstructure:"BEAST_LUCK_MAX"`     // maximum Beast luck
}

// SkillConfig have configuration for the stat Modifier
type SkillConfig struct {
	StrengthModifier float64 `mapstructure:"SKILL_DAMAGE_MODIFIER"`   // strength modifier for the skill
	StrengthChance   float64 `mapstructure:"SKILL_DAMAGE_CHANCE"`     // strength chance for the skill
	DamageModifier   float64 `mapstructure:"SKILL_STRENGTH_MODIFIER"` // Damage modifier for the skill
	DamageChance     float64 `mapstructure:"SKILL_STRENGTH_CHANCE"`   // Damage chance for the skill
}

// Config have configuration for the application
type Config struct {
	HeroConfig  `mapstructure:",squash"`
	BeastConfig `mapstructure:",squash"`
	SkillConfig `mapstructure:",squash"`
}

// Global contains the application configurations
var Global Config

// LoadConfig loads the config on the basis of the environment
// LoadConfig uses the following order to overide configurations
// 1. .default.env, 2. .$env.env, 3. environment variables
func LoadConfig(env, configFolderPath string) error {

	v := viper.New()
	v.SetTypeByDefaultValue(true)

	v.SetConfigName(".default")
	v.SetConfigType("env")
	v.AddConfigPath(configFolderPath)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if env != "" {
		env = strings.ToLower(env)
		v.SetConfigName("." + env)
		err := v.MergeInConfig()
		if err != nil {
			return err
		}
	}

	if err := v.Unmarshal(&Global); err != nil {
		return err
	}

	return nil
}
