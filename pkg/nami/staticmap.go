package nami

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type StaticMap []byte

const (
	staticmapURL = "https://maps.googleapis.com/maps/api/staticmap?"
)

func (d *Nami) FindStaticMap(origin, destination, apiKey string) (StaticMap, error) {
	key := generateStaticMapKey(origin, destination)
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

	staticmap, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	d.store.Set(key, staticmap)
	return staticmap, nil
}

func buildStaticMapURL(origin string, polyline DirectionPolyline, apiKey string) string {
	u := url.Values{}
	u.Add("center", origin)
	u.Add("size", "400x400")
	u.Add("path", fmt.Sprintf("weight:5|color:blue|enc:%s", polyline))
	u.Add("key", apiKey)
	return staticmapURL + u.Encode()
}
