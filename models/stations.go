package models

import (
	"github.com/carlqt/ez-bus/env"
	"github.com/carlqt/ez-bus/requests"
)

type Station struct {
	BusStopCode string  `db:"bus_stop_code"`
	RoadName    string  `db:"road_name"`
	Description string  `db:"description"`
	Latitude    float64 `db:"latitude"`
	Longitude   float64 `db:"longitude"`
	Buses       []Bus
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
	env.DBX.MustExec(q, s.BusStopCode, s.RoadName, s.Description, s.Latitude, s.Longitude)
}

// Nearby search the database for nearby stations within a given radius. The unit of radius is meters
func Nearby(radius int, c Location) (Stations, error) {
	var stations Stations

	err := env.DBX.Select(&stations, "SELECT description, bus_stop_code from stations WHERE earth_box(ll_to_earth($1, $2), $3) @> ll_to_earth(latitude, longitude)", c["lat"], c["lng"], radius)
	if err != nil {
		return nil, err
	}

	// stations.initBuses()
	return stations, nil
}

func (s *Stations) initBuses() error {
	stmnt, err := env.DBX.Preparex(`SELECT DISTINCT buses.bus_id_code
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
	err := env.DBX.Select(&s.Buses, `SELECT DISTINCT buses.bus_id_code FROM buses
	JOIN routes ON buses.bus_id_code = routes.bus_id_code
	WHERE routes.bus_stop_code = $1`, s.BusStopCode)

	if err != nil {
		return err
	}
	return nil
}

func (s *Station) GetStationBuses() error {
	busArrival := &requests.Bus{StationCode: s.BusStopCode}
	err := busArrival.Request()
	if err != nil {
		return err
	}

	for _, service := range busArrival.Services {
		bus := Bus{BusIDCode: service.ServiceNo}
		bus.Status = service.Status
		bus.NextBus = service.NextBus
		bus.SubsequentBus = service.SubsequentBus
		s.Buses = appendUniqueBus(s.Buses, bus)
		// appendUniqueBus(&s.Buses, bus)
	}

	return nil
}

func appendUniqueBus(dst []Bus, item Bus) []Bus {
	for _, bus := range dst {
		if bus.BusIDCode == item.BusIDCode {
			return dst
		}
	}

	return append(dst, item)
}
