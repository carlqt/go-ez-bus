package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"

	"github.com/carlqt/ez-bus/config"
	"github.com/carlqt/ez-bus/dbcon"
	"github.com/carlqt/ez-bus/models"
)

var db *sql.DB
var dbInfo string

func init() {
	var err error
	dbcon.DBcon, err = sql.Open("postgres", "dbname=sg_buses sslmode=disable")
	if err != nil {
		panic(err)
	}

	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(dbcon.DBcon)
	dbcon.SDBcon = &builder

}

func main() {
	busRouteRequest()
}

func busRouteRequest() {
	skipCtr := 0
	conf := config.NewConfig()
	url := conf.BaseURL + "ltaodataservice/" + "BusRoutes"

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
		routeResp := models.RouteResponse{}
		json.NewDecoder(resp.Body).Decode(&routeResp)

		routeResp.CreateAll()
		values = len(routeResp.Values)

		skipCtr += 50
		log.Println(url + "?$skip=" + strCtr)
	}
}

func debugResponse(r *http.Response) {
	var out bytes.Buffer
	body, _ := ioutil.ReadAll(r.Body)

	json.Indent(&out, body, "", "  ")
	out.WriteTo(os.Stdout)
}
