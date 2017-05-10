package models

import "github.com/carlqt/ez-bus/dbcon"

type RouteResponse struct {
	Values []Route `json:"Value"`
}

type Route struct {
	ServiceNo    string
	BusStopCode  string
	StopSequence int
}

func (r *Route) Create() {
	route := dbcon.SDBcon.Insert("routes").Columns("bus_stop_code", "bus_id_code", "stop_sequence").Values(r.BusStopCode, r.ServiceNo, r.StopSequence)
	_, err := route.Exec()

	if err != nil {
		panic(err)
	}
}

func (r *RouteResponse) CreateAll() {
	for _, route := range r.Values {
		route.Create()
		//fmt.Println(route)
	}
}
