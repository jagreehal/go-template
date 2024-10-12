package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Hello(name string) (string, error) {
	// Use viper to load configuration
	viper.SetDefault("greeting", "Hello")
	greeting := viper.GetString("greeting")

	// Use logrus for structured logging
	logrus.WithFields(logrus.Fields{
		"function": "Hello",
		"name":     name,
	}).Info("Hello function called")

	// Use pkg/errors for error wrapping (if needed)
	if name == "" {
		err := errors.New("name cannot be empty")
		// Wrap the error for more context
		return "", errors.Wrap(err, "invalid argument passed to Hello function")
	}

	return fmt.Sprintf("%s, %s!", greeting, name), nil
}

func main() {
	// Initialize logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Initialize viper (optional: load config from file)
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Fatal("Error reading config file")
	}

	result, err := Hello("World")
	if err != nil {
		logrus.WithError(err).Error("Error in Hello function")
		return
	}
	fmt.Println(result)
}
