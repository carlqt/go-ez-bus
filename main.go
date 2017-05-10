package main

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/carlqt/ez-bus/dbcon"
	"github.com/carlqt/ez-bus/models"
	_ "github.com/lib/pq"
)

func init() {
	var err error
	dbcon.DBcon, err = sql.Open("postgres", "dbname=sg_buses sslmode=disable")
	if err != nil {
		panic(err)
	}
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(dbcon.DBcon)
	dbcon.SDBcon = &builder
}

func main() {
	var stations models.Stations
	coords := make(map[string]float64)
	coords["lat"] = 1.293
	coords["lng"] = 103.8897

	stations, err := stations.Nearby(1000, coords)
	if err != nil {
		panic(err)
	}
	fmt.Println(stations)
}
