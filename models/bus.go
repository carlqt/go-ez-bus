package models

import (
	"github.com/carlqt/ez-bus/env"
	"github.com/carlqt/ez-bus/requests"
)

type Bus struct {
	BusIDCode     string `json:"ServiceNo" db:"bus_id_code"`
	BusOperator   string `json:"Operator" db:"bus_operator"`
	Status        string
	NextBus       requests.BusArrival
	SubsequentBus requests.BusArrival
}

//Create inserts bus record
func (b *Bus) Create() {
	q := "INSERT INTO buses (bus_id_code, bus_operator) VALUES ($1, $2)"
	env.DBX.MustExec(q, b.BusIDCode, b.BusOperator)
}

func (b *Bus) FetchArrival(stationCode string) error {
	busRequest := requests.Bus{BusCode: b.BusIDCode, StationCode: stationCode}
	err := busRequest.Request()

	services := busRequest.Services
	b.NextBus = services[0].NextBus
	b.SubsequentBus = services[0].SubsequentBus

	if err != nil {
		return err
	}

	return nil
}

func (b *Bus) FetchBusAndArrival(stationCode string) ([]*Bus, error) {
	var buses []*Bus

	err := env.DBX.Select(&buses, `SELECT DISTINCT buses.bus_id_code FROM buses
	JOIN routes ON buses.bus_id_code = routes.bus_id_code
	WHERE routes.bus_stop_code = $1`, stationCode)

	if err != nil {
		return nil, err
	}

	for _, bus := range buses {
		bus.FetchArrival(stationCode)
	}

	return buses, nil
}
