package server

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/cfx-cv/herald/pkg/common"
	"github.com/cfx-cv/nami/pkg/nami"
)

func (s *Server) staticmap(w http.ResponseWriter, r *http.Request) {
	d := nami.NewNami(s.store)
	origin, destination, apiKey := parseURL(r.URL.Query())

	result, err := d.FindStaticMap(origin, destination, apiKey)
	if err != nil {
		log.Print(err)
		common.Publish(common.NamiErrors, err.Error())
		return
	}

	data := map[string]interface{}{"staticmap": result}
	if err = json.NewEncoder(w).Encode(data); err != nil {
		log.Print(err)
		common.Publish(common.NamiErrors, err.Error())
		return
	}
}

func parseURL(u url.Values) (origin, destination, apiKey string) {
	origin = u.Get("origin")
	destination = u.Get("destination")
	apiKey = u.Get("api_key")
	return
}
