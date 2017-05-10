package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/carlqt/ez-bus/dbcon"
	_ "github.com/lib/pq"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func init() {
	var err error
	dbcon.DBcon, err = sql.Open("postgres", "dbname=sg_buses sslmode=disable")
	if err != nil {
		panic(err)
	}

	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(dbcon.DBcon)
	dbcon.SDBcon = &builder
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", Index)
	r.Get("/nearby", NearbyStations)
	log.Println("listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}