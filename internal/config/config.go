package config

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	Amqp struct {
		Dsn      string
		Exchange string
		Routing  struct {
			Vehicle string
			People  string
		}
	}
	Log struct {
		Level string
	}
	Listen struct {
		Host string
		Port int
	}
}

// NewConfig ...
func NewConfig() (*Config, error) {

	var level logrus.Level
	var err error
	config := &Config{}

	if err = viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file: %s", viper.ConfigFileUsed())
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal config")
	}

	level, err = logrus.ParseLevel(config.Log.Level)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal config")
	}

	logrus.SetLevel(level)
	//logrus.SetReportCaller(true)

	if level == logrus.DebugLevel {
		spew.Dump(config)
	}

	return config, nil
}
