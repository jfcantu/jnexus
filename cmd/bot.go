package cmd

import (
	"github.com/jfcantu/jnexus/pkg/bot"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Connects to the network and sends updates to the API service",
	Run: func(cmd *cobra.Command, args []string) {
		lvl, err := logrus.ParseLevel(viper.GetString("loglevel"))
		if err != nil {
			logrus.Warnf("Error parsing requested log level - defaulting to INFO: %v", err.Error())
		} else {
			logrus.SetLevel(lvl)
		}

		bot.Run()
	},
}

func init() {
	rootCmd.AddCommand(botCmd)

	botCmd.Flags().String("irc-server", "", "IRC server host/IP")
	botCmd.Flags().String("irc-server-port", "6667", "IRC server port")
	botCmd.Flags().Bool("irc-server-ssl", false, "Use SSL to connect to IRC server")
	botCmd.Flags().String("irc-server-password", "", "IRC server password")
	botCmd.Flags().String("irc-nick", "jnexus", "IRC nick")
	botCmd.Flags().String("irc-ident", "jnexus", "IRC ident")
	botCmd.Flags().String("irc-realname", "Construct additional pylons.", "IRC realname/GECOS")
	botCmd.Flags().String("irc-oper-username", "", "IRC oper username")
	botCmd.Flags().String("irc-oper-password", "", "IRC oper password")

	botCmd.Flags().String("api-server-addr", "127.0.0.1:7444", "ADDR:PORT of API server GRPC")

	viper.BindPFlags(botCmd.Flags())
}
