package bot

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Run starts the IRC bot
func Run() {
	b := NewBot()

	logrus.Infof("Connecting to %v", viper.GetString("irc-server"))
	b.IRCClient.Run()
}
