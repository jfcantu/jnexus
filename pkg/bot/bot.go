package bot

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/jfcantu/jnexus/pb"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
	"gopkg.in/irc.v3"
)

// Bot represents the IRC information-scraping agent
type Bot struct {
	Config     *Config
	IRCClient  *irc.Client
	GRPCClient pb.APIClient
}

// Config represents the bot's config (not including IRC client config)
type Config struct {
	OperUsername string
	OperPassword string
}

// NewBot creates a new Bot instance.
func NewBot() *Bot {
	b := new(Bot)

	// Connect to GRPC server
	grpcOptions := []grpc.DialOption{
		grpc.WithBlock(),
	}
	if viper.GetBool("allow-insecure-grpc") {
		grpcOptions = append(grpcOptions, grpc.WithInsecure())
	}
	if err := b.connectGRPC(grpcOptions...); err != nil {
		logrus.Fatal(err)
	}

	conn, err := net.Dial("tcp", fmt.Sprintf(
		"%s:%s",
		viper.GetString("irc-server"),
		viper.GetString("irc-server-port"),
	))
	if err != nil {
		logrus.Fatal(err)
	}

	b.IRCClient = irc.NewClient(conn, irc.ClientConfig{
		Nick:    viper.GetString("irc-nick"),
		Pass:    viper.GetString("irc-server-password"),
		User:    viper.GetString("irc-ident"),
		Name:    viper.GetString("irc-realname"),
		Handler: irc.HandlerFunc(b.MessageHandler),
	})

	return b
}

func (b *Bot) connectGRPC(opts ...grpc.DialOption) error {
	logrus.Infof("Connecting to GRPC server: %s", viper.GetString("api-server-addr"))

	dialCtx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()
	conn, err := grpc.DialContext(
		dialCtx,
		viper.GetString("api-server-addr"),
		opts...)
	if err != nil {
		return err
	}

	b.GRPCClient = pb.NewAPIClient(conn)

	return nil
}
