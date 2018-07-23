package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type Msg struct {
	Type    string `json:"type"`
	Path    string `json:"path"`
	Latency int64  `json:"latency"`
}

type Server struct{}

func (s *Server) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"request": r,
		}).Info("Received http request")

		if r.URL.Path != "/set_latency" {
			http.NotFound(w, r)
			return
		}

		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var msg Msg
		err = json.Unmarshal(b, &msg)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if msg.Type == "read" {
			globalInjuredHook.SetReadLatency(msg.Path, time.Duration(msg.Latency))
		} else if msg.Type == "fsync" {
			globalInjuredHook.SetFsyncLatency(msg.Path, time.Duration(msg.Latency))
		} else {
			http.Error(w, "Unsupport type", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
	})
}

func (s *Server) Run(addr string) error {
	httpServer := &http.Server{
		Addr:    addr,
		Handler: s.Handler(),
	}

	return httpServer.ListenAndServe()
}
