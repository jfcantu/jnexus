package bot

import (
	"context"
	"regexp"

	"github.com/jfcantu/jnexus/pb"
	"github.com/sirupsen/logrus"
	"gopkg.in/irc.v3"
)

const patternLinkClosedRemotely = `Routing -- from (?P<hub>[\w\.]+): Server (?P<leaf>[\w\.]+) closed the connection`
const patternLinkClosedLocally = `Routing -- from (?P<hub>[\w\.]+): No response from (?P<leaf>[\w\.]+),`
const patternLinkEstablished = `Routing -- from (?P<hub>[\w\.]+): Link with (?P<leaf>[\w\.]+) established`
const patternLinkCompleted = `Routing -- from (?P<hub>[\w\.]+): (?P<leaf>[\w\.]+) has synched to network data`

var regexLinkClosedLocally *regexp.Regexp = regexp.MustCompile(patternLinkClosedLocally)
var regexLinkClosedRemotely *regexp.Regexp = regexp.MustCompile(patternLinkClosedRemotely)
var regexLinkEstablished *regexp.Regexp = regexp.MustCompile(patternLinkEstablished)
var regexLinkCompleted *regexp.Regexp = regexp.MustCompile(patternLinkCompleted)

// RefreshMap just does /links
func (b *Bot) RefreshMap() {
	b.IRCClient.WriteMessage(&irc.Message{
		Command: "LINKS",
	})
}

// HandleRoutingMessage handles routing notices
func (b *Bot) HandleRoutingMessage(message string) {
	if match := regexLinkClosedLocally.FindStringSubmatch(message); match != nil {
		logrus.Infof("Received netsplit notification from %v for %v", match[1], match[2])
		b.GRPCClient.UpdateLinkStatus(
			context.Background(),
			&pb.LinkStatus{
				Server1: match[1],
				Server2: match[2],
				Status:  pb.LinkState_INACTIVE,
			},
		)
		return
	}

	if match := regexLinkClosedRemotely.FindStringSubmatch(message); match != nil {
		logrus.Infof("Received netsplit notification from %v for %v", match[1], match[2])
		b.GRPCClient.UpdateLinkStatus(
			context.Background(),
			&pb.LinkStatus{
				Server1: match[1],
				Server2: match[2],
				Status:  pb.LinkState_INACTIVE,
			},
		)
		return
	}

	if match := regexLinkEstablished.FindStringSubmatch(message); match != nil {
		logrus.Infof("Link established between %v and %v", match[1], match[2])
		b.GRPCClient.UpdateLinkStatus(
			context.Background(),
			&pb.LinkStatus{
				Server1: match[1],
				Server2: match[2],
				Status:  pb.LinkState_SYNCHRONIZING,
			},
		)
		return
	}

	if match := regexLinkCompleted.FindStringSubmatch(message); match != nil {
		logrus.Infof("Link between %v and %v is now synched", match[1], match[2])
		b.GRPCClient.UpdateLinkStatus(
			context.Background(),
			&pb.LinkStatus{
				Server1: match[1],
				Server2: match[2],
				Status:  pb.LinkState_ACTIVE,
			},
		)
		return
	}

	logrus.Info(message)
}
