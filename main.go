package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/carlqt/ez-bus/config"
)

func main() {
	conf := config.NewConfig()
	u := conf.BaseURL + "ltaodataservice/" + "BusRoutes"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", u, nil)
	//req.Header.Set("UniqueUserID", ltaUserKey)
	req.Header.Set("AccountKey", conf.BusKey)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	debugResponse(resp)

}

func debugResponse(r *http.Response) {
	var out bytes.Buffer
	body, _ := ioutil.ReadAll(r.Body)

	json.Indent(&out, body, "", "  ")
	out.WriteTo(os.Stdout)
}
