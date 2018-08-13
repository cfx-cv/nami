package nami

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type StaticMap []byte

const (
	staticmapURL = "https://maps.googleapis.com/maps/api/staticmap?"
)

func (d *Nami) FindStaticMap(origin, destination, apiKey string) (StaticMap, error) {
	key := generateKey(origin, destination)
	if value, ok := d.store.Get(key); ok {
		return StaticMap(value), nil
	}

	polyline, err := d.findDirectionPolyline(origin, destination, apiKey)
	if err != nil {
		return nil, err
	}

	url := buildStaticMapURL(origin, polyline, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var staticmap StaticMap
	if _, err := resp.Body.Read(staticmap); err != nil {
		log.Print(err)
		return nil, err
	}
	d.store.Set(key, staticmap)
	return staticmap, nil
}

func buildStaticMapURL(origin string, polyline DirectionPolyline, apiKey string) string {
	u := url.Values{}
	u.Add("center", origin)
	u.Add("size", "size=400x400")
	u.Add("path", fmt.Sprintf("weight:5|color:blue|enc:%s", polyline))
	u.Add("key", apiKey)
	return staticmapURL + u.Encode()
}
