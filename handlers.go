package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/carlqt/ez-bus/models"
	"github.com/pressly/chi"
)

func ApplicationHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		next.ServeHTTP(w, r)
	})
}

func BusStopAuth(next http.HandlerFunc) http.HandlerFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func NearbyStations(w http.ResponseWriter, r *http.Request) {
	var stationsStruct models.Stations
	var resp []byte
	coors := make(map[string]float64)

	q := r.URL.Query()
	coors["lat"], _ = strconv.ParseFloat(q.Get("lat"), 64)
	coors["lng"], _ = strconv.ParseFloat(q.Get("lng"), 64)

	stations, err := stationsStruct.Nearby(1000, coors)
	if err != nil {
		errMap := make(map[string]string)
		errMap["errors"] = err.Error()
		resp, _ = json.Marshal(errMap)
		w.Write(resp)
	} else {
		resp, _ := json.Marshal(stations)
		fmt.Println(stations)
		w.Write(resp)
	}
}

func Station(w http.ResponseWriter, r *http.Request) {
	var stationList models.Stations
	var resp []byte

	q := r.URL.Query()
	busCode := chi.URLParam(r, "busStopCode")
	station := q.Get("stationCode")

	stations, err := stationList.RemainingRoute(busCode, station)
	if err != nil {
		errMap := make(map[string]string)
		errMap["errors"] = err.Error()
		resp, _ = json.Marshal(errMap)
		w.Write(resp)
	} else {
		resp, _ = json.Marshal(stations)
		w.Write(resp)
	}
}
