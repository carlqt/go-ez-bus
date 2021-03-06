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

type StationResponse struct {
	Values []models.Station `json:"Value"`
}

func (b *StationResponse) CreateAll() {
	for _, station := range b.Values {
		if !StationExists(station.BusStopCode) {
			station.Create()
		}
	}
}

func StationExists(code string) bool {
	var exists bool
	env.DBX.QueryRowx("SELECT exists (SELECT 1 FROM stations WHERE bus_stop_code = $1)", code).Scan(&exists)

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
	url := conf.BaseURL + "ltaodataservice/" + "BusStops"

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	for values := 50; values >= 50; {
		strCtr := strconv.Itoa(skipCtr)
		req, _ := http.NewRequest("GET", url+"?$skip="+strCtr, nil)
		req.Header.Set("AccountKey", conf.BusKey)
		req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		stationResp := StationResponse{}
		json.NewDecoder(resp.Body).Decode(&stationResp)

		stationResp.CreateAll()
		values = len(stationResp.Values)

		skipCtr += 50
		log.Println(url + "?$skip=" + strCtr)
	}
}
