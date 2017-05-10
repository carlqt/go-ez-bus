package models

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/carlqt/ez-bus/dbcon"
)

type StationResponse struct {
	Values Stations `json:"Value"`
}

type Station struct {
	BusStopCode string
	RoadName    string
	Description string
	Latitude    float32
	Longitude   float32
}

type Stations []Station

type Location map[string]float64

func (s *Station) Create() {
	station := dbcon.SDBcon.Insert("stations").Columns("bus_stop_code", "road_name", "description", "latitude", "longitude").Values(s.BusStopCode, s.RoadName, s.Description, s.Latitude, s.Longitude)
	_, err := station.Exec()

	if err != nil {
		panic(err)
	}
}

func (b *StationResponse) CreateAll() {
	for _, station := range b.Values {
		if !StationExists(station.BusStopCode) {
			station.Create()
		}
	}
}

// Nearby search the database for nearby stations within a given radius
func (s *Stations) Nearby(radius int, c Location) (Stations, error) {
	var stations []Station

	qBuilder := dbcon.SDBcon.Select("description").From("stations").Where("earth_box(ll_to_earth($1, $2), $3) @> ll_to_earth(latitude, longitude)", c["lat"], c["lng"], radius)

	rows, err := qBuilder.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		station := Station{}
		rows.Scan(&station.Description)
		stations = append(stations, station)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stations, nil
}

func (s *Stations) RemainingRoute(busCode string, stationCode string) (Stations, error) {
	var stopCount int
	var stations []Station

	err := dbcon.DBcon.QueryRow(`SELECT stop_sequence FROM routes
	WHERE bus_id_code = $1 AND bus_stop_code = $2`, busCode, stationCode).Scan(&stopCount)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, err := dbcon.DBcon.Query(`SELECT stations.description, stations.bus_stop_code FROM stations
	JOIN routes ON stations.bus_stop_code = routes.bus_stop_code
	WHERE routes.bus_id_code = $1 AND routes.stop_sequence > $2
	ORDER BY routes.stop_sequence`, busCode, stopCount)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		station := Station{}
		rows.Scan(&station.Description, &station.BusStopCode)
		stations = append(stations, station)
	}

	return stations, nil
}

func StationExists(code string) bool {
	var stnCode string
	err := dbcon.SDBcon.Select("bus_stop_code").From("stations").Where(sq.Eq{"bus_stop_code": code}).QueryRow().Scan(&stnCode)

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	default:
		return true
	}
}
