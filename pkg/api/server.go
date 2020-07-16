package api

import (
	"net/http"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func newServer() *Server {
	s := new(Server)
	s.WaitGroup = new(sync.WaitGroup)

	return s
}

func (s *Server) serveAPI() {
	logrus.Infof("Starting HTTP listener on: %v", viper.GetString("api-server-bind-addr"))

	m := mux.NewRouter()
	m.HandleFunc("/networkstatus", s.getNetworkStatus)
	m.HandleFunc("/servers", s.getServers)
	m.HandleFunc("/links/active", s.getActiveLinks)
	m.HandleFunc("/links/all", s.getAllLinks)
	m.HandleFunc("/servers/split", s.getSplitServers)

	s.HTTPServer = &http.Server{
		Handler: handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"POST", "OPTIONS", "DELETE", "PUT"}),
			handlers.AllowedHeaders([]string{"content-type"}),
		)(m),
		Addr: viper.GetString("api-server-bind-addr"),
	}

	s.WaitGroup.Add(1)

	// Another blocking call, so start a goroutine that decrements the waitgroup when it finishes
	go func() {
		defer s.WaitGroup.Done()
		if err := s.HTTPServer.ListenAndServe(); err != http.ErrServerClosed {
			logrus.Fatal(err.Error())
		}
	}()
}
