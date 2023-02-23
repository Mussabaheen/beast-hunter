// Package main represents the entry point of the game
package main

import (
	"os"
	"strings"

	"github.com/mussabaheen/beast-hunter/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {

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
}
