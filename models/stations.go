package models

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/carlqt/ez-bus/dbcon"
)

type StationResponse struct {
	Values []Station `json:"Value"`
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

func (s *Station) Create(db *sql.DB) {
	sgBusDB := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db)
	station := sgBusDB.Insert("stations").Columns("bus_stop_code", "road_name", "description", "latitude", "longitude").Values(s.BusStopCode, s.RoadName, s.Description, s.Latitude, s.Longitude)
	_, err := station.Exec()

	if err != nil {
		panic(err)
	}
}

func (b *StationResponse) CreateAll(db *sql.DB) {
	for _, station := range b.Values {
		if !Exists(db, station.BusStopCode) {
			station.Create(db)
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

func StationExists(db *sql.DB, code string) bool {
	var stnCode string
	sgBusDB := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db)
	err := sgBusDB.Select("bus_stop_code").From("stations").Where(sq.Eq{"bus_stop_code": code}).QueryRow().Scan(&stnCode)

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	default:
		return true
	}
}
