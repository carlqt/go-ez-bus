package geocode

import (
	"log"
	"net/url"
	"os"
	"strconv"
)

var key string
var secret string
var baseURL = "http://www.mapquestapi.com/geocoding/v1/reverse"

type location map[string]string

// InitData Initialize app key and code
func InitData(k string, s string) {
	key = k
	secret = s
}

func Reverse(lat float64, lng float64) {
	coors := make(location)
	coors["lat"] = strconv.FormatFloat(lat, 'f', -1, 64)
	coors["lng"] = strconv.FormatFloat(lng, 'f', -1, 64)

	u := urlBuilder(coors)
	geoRevRequest(u)
}

func urlBuilder(coors location) url.URL {
	u, _ := url.Parse(baseURL)
	q := make(url.Values)
	q.Add("key", os.Getenv("MQ_API_KEY"))
	q.Add("location", coors["lat"]+","+coors["lng"])

	u.RawQuery = q.Encode()
	return *u
}
