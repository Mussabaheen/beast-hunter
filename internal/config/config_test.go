package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_readsDefault(t *testing.T) {
	configDir := t.TempDir()
	conf := "HERO_HEALTH_MIN=70\nHERO_HEALTH_MAX=100"
	err := os.WriteFile(fmt.Sprintf("%s/.default.env", configDir), []byte(conf), 0644)

	assert.Nil(t, err, "error must be Nil")

	err = LoadConfig("", configDir)
	assert.Nil(t, err, "error must be Nil")

	assert.Equal(t, 70, Global.HeroConfig.HealthMin)
	assert.Equal(t, 100, Global.HeroConfig.HealthMax)
}

func TestLoadConfig_errorConfigDirectoryNotProvided(t *testing.T) {
	err := LoadConfig("", "")
	assert.NotNil(t, err)
}

func TestLoadConfig_overridesDefaultWithEnvFile(t *testing.T) {
	configDir := t.TempDir()
	defaultConf := "HERO_HEALTH_MIN=70\nHERO_HEALTH_MAX=100"
	prodConf := "HERO_HEALTH_MIN=40\nHERO_HEALTH_MAX=60"

	err := os.WriteFile(fmt.Sprintf("%s/.default.env", configDir), []byte(defaultConf), 0644)
	assert.Nil(t, err, "error must be Nil")

	err = os.WriteFile(fmt.Sprintf("%s/.production.env", configDir), []byte(prodConf), 0644)
	assert.Nil(t, err, "error must be Nil")

	err = LoadConfig("production", configDir)
	assert.Nil(t, err, "error must be Nil")
	assert.Equal(t, 40, Global.HeroConfig.HealthMin)
	assert.Equal(t, 60, Global.HeroConfig.HealthMax)
}
func TestLoadConfig_overridesDefaultWithEnvironmentVariables(t *testing.T) {
	configDir := t.TempDir()

	conf := "HERO_HEALTH_MIN=70\nHERO_HEALTH_MAX=80"
	err := os.WriteFile(fmt.Sprintf("%s/.default.env", configDir), []byte(conf), 0644)
	assert.Nil(t, err, "error must be Nil")

	err = os.Setenv("HERO_HEALTH_MIN", "90")
	assert.Nil(t, err, "error must be Nil")

	err = os.Setenv("HERO_HEALTH_MAX", "100")
	assert.Nil(t, err, "error must be Nil")

	err = LoadConfig("", configDir)
	assert.Nil(t, err, "error must be Nil")

	assert.Equal(t, 90, Global.HeroConfig.HealthMin)
	assert.Equal(t, 100, Global.HeroConfig.HealthMax)
}
