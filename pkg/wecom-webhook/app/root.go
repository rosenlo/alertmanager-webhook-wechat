package app

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Version   = "Unkown"
	GitHash   = "Unkown"
	BuildTime = "Unkown"
)

// weComWebhook represents the base command when called without any subcommands
var weComWebhook = &cobra.Command{
	Use:     "wecom-webhook",
	Short:   "WeCom Webhook for Alertmanager",
	Long:    `WeCom Webhook for Alertmanager`,
	Version: fmt.Sprintf("%s-%s-%s", Version, GitHash, BuildTime),
	Run:     Run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return weComWebhook.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix("WECOM_WEBHOOK")
	// basic
	viper.SetDefault("log_level", "debug")
	viper.SetDefault("address", ":9090")
	viper.SetDefault("robot_url", "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=$key")
	viper.SetDefault("user_url", "http://example.com/weixin/send")

	viper.AutomaticEnv() // read in environment variables that match

	// test config
	if viper.GetString("env") == "test" {
	}
}
