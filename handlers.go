package main

import (
	"net/http"
	"strconv"

	"github.com/carlqt/ez-bus/helpers/render"
	"github.com/carlqt/ez-bus/models"
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
	h := make(map[string]interface{})
	coors := make(map[string]float64)

	q := r.URL.Query()
	coors["lat"], _ = strconv.ParseFloat(q.Get("lat"), 64)
	coors["lng"], _ = strconv.ParseFloat(q.Get("lng"), 64)

	stations, err := models.Nearby(1000, coors)
	h["stations"] = stations
	h["errors"] = err
	render.JSON(w, h)
}

func Station(w http.ResponseWriter, r *http.Request) {
	// var stationList models.Stations
	// h := make(map[string]interface{})

	// q := r.URL.Query()
	// busCode := chi.URLParam(r, "busStopCode")
	// station := q.Get("stationCode")

	// stations, err := stationList.RemainingRoute(busCode, station)
	// h["stations"] = stations
	// h["errors"] = err
	// render.JSON(w, h)
}
