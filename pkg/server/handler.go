package server

import (
	"log"
	"net/http"
	"net/url"

	"github.com/cfx-cv/nami/pkg/nami"
)

func (s *Server) staticmap(w http.ResponseWriter, r *http.Request) {
	d := nami.NewNami(s.store)
	origin, destination, apiKey := parseURL(r.URL.Query())

	result, err := d.FindStaticMap(origin, destination, apiKey)
	if err != nil {
		log.Print(err)
		return
	}

	if _, err = w.Write(result); err != nil {
		log.Print(err)
		return
	}
}

func parseURL(u url.Values) (origin, destination, apiKey string) {
	origin = u.Get("origin")
	destination = u.Get("destination")
	apiKey = u.Get("api_key")
	return
}
