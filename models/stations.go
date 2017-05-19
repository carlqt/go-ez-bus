package models

import (
	"github.com/carlqt/ez-bus/dbcon"
)

type Station struct {
	BusStopCode string  `db:"bus_stop_code"`
	RoadName    string  `db:"road_name"`
	Description string  `db:"description"`
	Latitude    float64 `db:"latitude"`
	Longitude   float64 `db:"longitude"`
	Buses       []string
}

type Location map[string]float64
type Stations []*Station

func NewStation(code string) *Station {
	station := &Station{BusStopCode: code}

	station.InitBuses()
	return station
}

func (s *Station) Create() {
	q := `INSERT INTO stations (bus_stop_code, road_name, description, latitude, longitude)
	VALUES ($1, $2, $3, $4, $5)`
	dbcon.DBX.MustExec(q, s.BusStopCode, s.RoadName, s.Description, s.Latitude, s.Longitude)
}

// Nearby search the database for nearby stations within a given radius. The unit of radius is meters
func Nearby(radius int, c Location) (Stations, error) {
	var stations Stations

	err := dbcon.DBX.Select(&stations, "SELECT description, bus_stop_code from stations WHERE earth_box(ll_to_earth($1, $2), $3) @> ll_to_earth(latitude, longitude)", c["lat"], c["lng"], radius)
	if err != nil {
		return nil, err
	}

	stations.initBuses()
	return stations, nil
}

// RemainingRoute returns a list of stations. If bus is given, it will return the list of incoming stations along the route of the bus
// func (s *Stations) RemainingRoute(busCode string, stationCode string) (Stations, error) {
// 	var stopCount int
// 	var stations []*Station

// 	// Get the stopcount of stations of the bus
// 	err := dbcon.DBcon.QueryRow(`SELECT stop_sequence FROM routes
// 	WHERE bus_id_code = $1 AND bus_stop_code = $2`, busCode, stationCode).Scan(&stopCount)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}

// 	rows, err := dbcon.DBcon.Query(`SELECT stations.description, stations.bus_stop_code FROM stations
// 	JOIN routes ON stations.bus_stop_code = routes.bus_stop_code
// 	WHERE routes.bus_id_code = $1 AND routes.stop_sequence > $2
// 	ORDER BY routes.stop_sequence`, busCode, stopCount)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		station := new(Station)
// 		rows.Scan(&station.Description, &station.BusStopCode)
// 		stations = append(stations, station)
// 	}

// 	return stations, nil
// }

func (s *Stations) initBuses() error {
	stmnt, err := dbcon.DBX.Preparex(`SELECT DISTINCT buses.bus_id_code
	FROM buses JOIN routes ON buses.bus_id_code = routes.bus_id_code
	WHERE routes.bus_stop_code = $1`)

	if err != nil {
		panic(err)
	}

	for _, station := range *s {
		stmnt.Select(&station.Buses, station.BusStopCode)
	}

	return nil
}

func (s *Station) InitBuses() error {
	err := dbcon.DBX.Select(&s.Buses, `SELECT DISTINCT buses.bus_id_code FROM buses
	JOIN routes ON buses.bus_id_code = routes.bus_id_code
	WHERE routes.bus_stop_code = $1`, s.BusStopCode)

	if err != nil {
		return err
	}
	return nil
}
