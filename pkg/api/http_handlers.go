package api

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *Server) getNetworkStatus(w http.ResponseWriter, r *http.Request) {
	logrus.Debug(r)
	network, err := s.DBClient.GetNetworkStatus()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	b, err := json.Marshal(network)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(200)
	w.Write(b)
}

func (s *Server) getSplitServers(w http.ResponseWriter, r *http.Request) {
	logrus.Debug(r)

	nodeList, err := s.DBClient.GetSplitServers()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	idList := []int{}

	for _, v := range nodeList {
		idList = append(idList, v.ID)
	}

	b, err := json.Marshal(idList)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(200)
	w.Write(b)
}

func (s *Server) getActiveLinks(w http.ResponseWriter, r *http.Request) {
	logrus.Debug(r)

	links, err := s.DBClient.GetActiveLinks()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	b, err := json.Marshal(links)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(200)
	w.Write(b)
}

func (s *Server) getAllLinks(w http.ResponseWriter, r *http.Request) {
	logrus.Debug(r)

	links, err := s.DBClient.GetAllLinks()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	b, err := json.Marshal(links)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(200)
	w.Write(b)
}

func (s *Server) getServers(w http.ResponseWriter, r *http.Request) {
	logrus.Debug(r)

	nodes, err := s.DBClient.GetServers()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	b, err := json.Marshal(nodes)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(200)
	w.Write(b)
}
