package models

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
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
