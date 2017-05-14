package models

import (
	"github.com/carlqt/ez-bus/dbcon"
)

type Bus struct {
	BusIDCode   string `json:"ServiceNo" db:"bus_id_code"`
	BusOperator string `json:"Operator" db:"bus_operator"`
}

//Create inserts bus record
func (b *Bus) Create() {
	q := "INSERT INTO buses (bus_id_code, bus_operator) VALUES ($1, $2)"
	dbcon.DBX.MustExec(q, b.BusIDCode, b.BusOperator)
}
