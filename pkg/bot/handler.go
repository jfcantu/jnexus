package bot

import (
	"context"
	"regexp"

	"github.com/jfcantu/jnexus/pb"
	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
	"gopkg.in/irc.v3"
)

var regexRoutingNotice *regexp.Regexp = regexp.MustCompile(`\*\*\* Routing`)

// MessageHandler handles messages from the server
func (b *Bot) MessageHandler(cli *irc.Client, m *irc.Message) {
	logrus.Debug(m)
	switch cmd := m.Command; cmd {
	case "001":
		cli.WriteMessage(&irc.Message{
			Command: "MODE",
			Params: []string{
				viper.GetString("irc-nick"),
				"+F",
			},
		})
		cli.WriteMessage(&irc.Message{
			Command: "OPER",
			Params: []string{
				viper.GetString("irc-oper-username"),
				viper.GetString("irc-oper-password"),
			},
		})
	case "381":
		b.RefreshMap()
	case "364":
		b.GRPCClient.UpdateLinkStatus(
			context.Background(),
			&pb.LinkStatus{
				Server1: m.Params[1],
				Server2: m.Params[2],
				Status:  pb.LinkState_ACTIVE,
			},
		)
	case "NOTICE":
		if regexRoutingNotice.MatchString(m.Params[1]) {
			b.HandleRoutingMessage(m.Params[1])
		}
	}
}
