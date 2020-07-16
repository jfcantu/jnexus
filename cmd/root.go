package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("config", "p", "", "Configuration file name.")
	rootCmd.PersistentFlags().String("loglevel", "INFO", "Log level.")

	rootCmd.PersistentFlags().Bool("allow-insecure-grpc", false, "Allow insecure gRPC transport.")
}

var rootCmd = &cobra.Command{
	Use:   "jnexus",
	Short: "it's like rnexus, but dire",
}

// Execute runs the bot
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logrus.Fatal(err)
		}
	}

	viper.SetEnvPrefix("JNEXUS")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}
