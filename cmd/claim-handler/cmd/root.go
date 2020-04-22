package cmd

import (
	"github.com/alexey-zayats/claim-handler/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "claim-handler",
	Short: "handle claims",
	Long:  "handle from claims for passes",
	Run:   rootMain,
}

func rootMain(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		logrus.WithFields(logrus.Fields{"reason": err}).Fatal("unable call cmd.Help")
	}
}

// ----------

// Execute entry point to the app
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.WithFields(logrus.Fields{"reason": err}).Fatal("unable execute rootCmd")
	}
}

func init() {

	viper.SetEnvPrefix(config.EnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	cfgParams := []config.Param{
		{Name: "log-level", Value: "info", Usage: "log level", ViperBind: "Log.Level"},
		{Name: "log-caller", Value: false, Usage: "log caller", ViperBind: "Log.Caller"},

		{Name: "listen-host", Value: "0.0.0.0", Usage: "listen host", ViperBind: "Listen.Host"},
		{Name: "listen-port", Value: 8080, Usage: "listen port", ViperBind: "Listen.Port"},

		{Name: "amqp-dsn", Value: "amqp://pass:pass@127.0.0.1:5672/", Usage: "AMQP datasource", ViperBind: "Amqp.Dsn"},
		{Name: "amqp-exchange", Value: "collector", Usage: "AMQP Exchange name publish to", ViperBind: "Amqp.Exchange"},

		{Name: "amqp-routing-vehicle", Value: "form.vehicle", Usage: "vehicle form routing key", ViperBind: "Amqp.Routing.Vehicle"},
		{Name: "amqp-routing-people", Value: "form.people", Usage: "people form routing key", ViperBind: "Amqp.Routing.People"},

		{Name: "cache-expire", Value: 1, Usage: "cache expire time", ViperBind: "Cache.Expire"},
		{Name: "cache-cleanup", Value: 5, Usage: "cache cleanup time", ViperBind: "Cache.Cleanup"},
	}

	config.Apply(rootCmd, cfgParams)
	viper.AutomaticEnv()

	rootCmd.PersistentFlags().StringVar(&config.FilePath, "config", "config.yaml", "Config file")

	cobra.OnInitialize(config.Init)
}
