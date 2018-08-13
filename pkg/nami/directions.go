package nami

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	directionURL = "https://maps.googleapis.com/maps/api/directions/json?"
)

type DirectionPolyline string

func (d *Nami) findDirectionPolyline(origin, destination, apiKey string) (DirectionPolyline, error) {
	key := generateDirectionKey(origin, destination)
	if value, ok := d.store.Get(key); ok {
		return DirectionPolyline(value), nil
	}

	url := buildDirectionURL(origin, destination, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := parseDirectionResponse(resp)
	if err != nil {
		return "", err
	}
	defer d.store.Set(key, []byte(result))
	return result, nil
}

func buildDirectionURL(origin, destination, apiKey string) string {
	u := url.Values{}
	u.Add("origin", origin)
	u.Add("destination", destination)
	u.Add("key", apiKey)
	return directionURL + u.Encode()
}

func parseDirectionResponse(resp *http.Response) (DirectionPolyline, error) {
	var body struct {
		Routes []struct {
			OverviewPolyline struct {
				Points DirectionPolyline `json:"points"`
			} `json:"overview_polyline"`
		} `json:"routes"`

		Status string `json:"status"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return "", err
	}
	if status := body.Status; status != "OK" {
		return "", errors.New(status)
	}

	return body.Routes[0].OverviewPolyline.Points, nil
}
