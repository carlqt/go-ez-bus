package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/carlqt/ez-bus/config"
	"github.com/carlqt/ez-bus/env"
	"github.com/carlqt/ez-bus/models"
)

type BusResponse struct {
	Values []models.Bus `json:"Value"`
}

func (b *BusResponse) CreateAll() {
	for _, bus := range b.Values {
		if !BusExists(bus.BusIDCode) {
			bus.Create()
		}
	}
}

func BusExists(busIDCode string) bool {
	var exists bool
	env.DBX.QueryRowx("SELECT exists (SELECT 1 FROM buses WHERE bus_id_code = $1)", busIDCode).Scan(&exists)

	switch {
	case exists:
		return true
	default:
		return false
	}
}

func init() {
	var err error
	env.DBX, err = sqlx.Connect("postgres", "dbname=sg_buses sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func main() {
	busRouteRequest()
}

func busRouteRequest() {
	skipCtr := 0
	conf := config.NewConfig()
	url := conf.BaseURL + "ltaodataservice/" + "BusServices"

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	for values := 50; values >= 50; {
		strCtr := strconv.Itoa(skipCtr)
		req, _ := http.NewRequest("GET", url+"?$skip="+strCtr, nil)
		//req.Header.Set("UniqueUserID", ltaUserKey)
		req.Header.Set("AccountKey", conf.BusKey)
		req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		busResp := BusResponse{}
		json.NewDecoder(resp.Body).Decode(&busResp)

		busResp.CreateAll()
		values = len(busResp.Values)

		skipCtr += 50
		log.Println(url + "?$skip=" + strCtr)
	}
}
