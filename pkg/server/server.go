package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"

	"github.com/cfx-cv/nami/pkg/nami"
	dredis "github.com/cfx-cv/nami/pkg/redis"
)

const (
	staticmapURL string = "/staticmap"
)

type Server struct {
	store nami.Store
}

func NewServer(client *redis.Client, expiration time.Duration) *Server {
	store := dredis.NewStore(client, expiration)
	return &Server{store: store}
}

func (s *Server) Start() {
	router := mux.NewRouter()
	router.HandleFunc(staticmapURL, s.staticmap).Methods("GET")

	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal(err)
	}
}
