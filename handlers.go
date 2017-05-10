package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/carlqt/ez-bus/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func NearbyStations(w http.ResponseWriter, r *http.Request) {
	var stationsStruct models.Stations
	coors := make(map[string]float64)

	q := r.URL.Query()
	coors["lat"], _ = strconv.ParseFloat(q.Get("lat"), 64)
	coors["lng"], _ = strconv.ParseFloat(q.Get("lng"), 64)

	stations, err := stationsStruct.Nearby(1000, coors)
	logErr(err)

	resp, _ := json.Marshal(stations)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func BusRoute(w http.ResponseWriter, r *http.Request) {
	var stationList models.Stations
	q := r.URL.Query()
	bus := q.Get("busCode")
	station := q.Get("stationCode")

	stations, err := stationList.RemainingRoute(bus, station)
	logErr(err)

	resp, _ := json.Marshal(stations)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
