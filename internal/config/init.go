package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var (
	// FilePath ...
	FilePath string

	// EnvPrefix ...
	EnvPrefix = "HANDLER"
)

// Init ...
func Init() {
	logrus.SetLevel(logrus.TraceLevel)

	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stderr)

	viper.SetConfigFile(FilePath)
}
