package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/carlqt/ez-bus/config"
	"github.com/carlqt/ez-bus/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func init() {
	var err error

	env.DBX, err = sqlx.Connect("postgres", "dbname=sg_buses sslmode=disable user=carlqt password=password")
	if err != nil {
		panic(err)
	}

	env.Conf = config.NewConfig()

	// Use this client on http requests to take advantage of keep-alive connections
	tr := new(http.Transport)
	env.HttpClient = &http.Client{Transport: tr}
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "build")
	r.FileServer("/", http.Dir(filesDir))
	r.Get("/stations", serveIndex)

	r.Route("/api", func(r chi.Router) {
		r.Use(ApiHandler)
		r.Get("/nearby", NearbyStations)
		r.Get("/station/:busStopCode", BusStopAuth(Station))
		r.Get("/station/:busStopCode/arrivals", stationBusArrival)
		r.Get("/stations", stations)
	})

	log.Println("listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
