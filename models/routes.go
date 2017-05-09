package models

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type RouteResponse struct {
	Values []Route `json:"Value"`
}

type Route struct {
	ServiceNo    string
	BusStopCode  string
	StopSequence int
}

func (r *Route) Create(db *sql.DB) {
	sgBusDB := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db)
	route := sgBusDB.Insert("routes").Columns("bus_stop_code", "bus_id_code", "stop_sequence").Values(r.BusStopCode, r.ServiceNo, r.StopSequence)
	_, err := route.Exec()

	if err != nil {
		panic(err)
	}
}

func (r *RouteResponse) CreateAll(db *sql.DB) {
	for _, route := range r.Values {
		route.Create(db)
		//fmt.Println(route)
	}
}
