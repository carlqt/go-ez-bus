package models

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type BusResponse struct {
	Values []Bus `json:"Value"`
}

type Bus struct {
	BusIDCode   string `json:"ServiceNo"`
	BusOperator string `json:"Operator"`
}

func (b *Bus) Create(db *sql.DB) {
	sgBusDB := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db)
	bus := sgBusDB.Insert("buses").Columns("bus_id_code", "bus_operator").Values(b.BusIDCode, b.BusOperator)
	_, err := bus.Exec()

	if err != nil {
		panic(err)
	}
}

func (b *BusResponse) CreateAll(db *sql.DB) {
	for _, bus := range b.Values {
		if !Exists(db, bus.BusIDCode) {
			bus.Create(db)
		}
	}
}

func Exists(db *sql.DB, busIDCode string) bool {
	var busID string
	sgBusDB := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db)
	err := sgBusDB.Select("bus_id_code").From("buses").Where(sq.Eq{"bus_id_code": busIDCode}).QueryRow().Scan(&busID)

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	default:
		return true
	}
}
