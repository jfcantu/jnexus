package api

import (
	"net/http"
	"sync"

	"google.golang.org/grpc"

	"github.com/jfcantu/jnexus/pb"

	"github.com/jfcantu/jnexus/pkg/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Server represents the API server
type Server struct {
	pb.APIServer

	WaitGroup *sync.WaitGroup

	DBClient   *db.Client
	HTTPServer *http.Server
	GRPCServer *grpc.Server
}

// Run starts the API service
func Run() {
	// Create a new Server object
	s := newServer()

	logrus.Infof("Connecting to database at: %v", viper.GetString("db-server-string"))
	// Connect to database
	dbClient, err := db.New(&db.Config{
		ServerString: viper.GetString("db-server-string"),
		Username:     viper.GetString("db-server-username"),
		Password:     viper.GetString("db-server-password"),
	})
	if err != nil {
		logrus.Fatal(err.Error())
	}

	defer func() {
		(*dbClient.Session).Close()
		(*dbClient.Driver).Close()
	}()

	s.DBClient = dbClient

	if err := s.DBClient.ClearRouteStates(); err != nil {
		logrus.Warnf("Error clearing route states - old route data may persist: %v", err)
	}

	// Start GRPC server
	s.serveGRPC()

	// Start HTTP server
	s.serveAPI()

	s.WaitGroup.Wait()
}
