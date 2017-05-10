package models

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/carlqt/ez-bus/dbcon"
)

type BusResponse struct {
	Values []Bus `json:"Value"`
}

type Bus struct {
	BusIDCode   string `json:"ServiceNo"`
	BusOperator string `json:"Operator"`
}

func (b *Bus) Create() {
	bus := dbcon.SDBcon.Insert("buses").Columns("bus_id_code", "bus_operator").Values(b.BusIDCode, b.BusOperator)
	_, err := bus.Exec()

	if err != nil {
		panic(err)
	}
}

func (b *BusResponse) CreateAll() {
	for _, bus := range b.Values {
		if !Exists(bus.BusIDCode) {
			bus.Create()
		}
	}
}

func Exists(busIDCode string) bool {
	var busID string
	err := dbcon.SDBcon.Select("bus_id_code").From("buses").Where(sq.Eq{"bus_id_code": busIDCode}).QueryRow().Scan(&busID)

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	default:
		return true
	}
}
