package cmd

import (
	"github.com/jfcantu/jnexus/pkg/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Accepts updates from the bot",
	Run: func(cmd *cobra.Command, args []string) {
		lvl, err := logrus.ParseLevel(viper.GetString("loglevel"))
		if err != nil {
			logrus.Warnf("Error parsing requested log level - defaulting to INFO: %v", err.Error())
		} else {
			logrus.Infof("Set log level to: %v", lvl)
			logrus.SetLevel(lvl)
		}

		api.Run()
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)

	serviceCmd.Flags().String("db-server-string", "bolt://localhost:7687", "neo4j server string")
	serviceCmd.Flags().String("db-server-username", "", "neo4j server username")
	serviceCmd.Flags().String("db-server-password", "", "neo4j server password")

	serviceCmd.Flags().String("api-server-bind-addr", "", "addr:port for API server to listen on")
	serviceCmd.Flags().String("grpc-server-bind-addr", "", "addr:port for GRPC server to listen on")

	viper.BindPFlags(serviceCmd.Flags())
}
