package main

import (
	"log"
	"net/http"

	"github.com/carlqt/ez-bus/dbcon"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func init() {
	var err error

	dbcon.DBcon, err = sqlx.Connect("postgres", "dbname=sg_buses sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(ApplicationHandler)
	r.Get("/", Index)
	r.Get("/nearby", NearbyStations)
	r.Get("/station/:busStopCode", BusStopAuth(Station))
	log.Println("listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
