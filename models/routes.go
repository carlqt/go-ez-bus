package models

import (
	"github.com/carlqt/ez-bus/dbcon"
)

type Route struct {
	ServiceNo    string
	BusStopCode  string
	StopSequence int
}

func (r *Route) Create() {
	q := "INSERT INTO routes (bus_stop_code, bus_id_code, stop_sequence) VALUES ($1, $2, $3)"
	dbcon.DBX.MustExec(q, r.BusStopCode, r.ServiceNo, r.StopSequence)
}
