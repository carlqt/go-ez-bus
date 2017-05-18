package requests

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/carlqt/ez-bus/dbcon"
)

type Bus struct {
	BusCode     string
	StationCode string `json:"BusStopID"`
	Services    []Service
}

type Service struct {
	ServiceNo     string
	Status        string
	OriginatingID string
	NextBus       BusArrival
	SubsequentBus BusArrival
}

type BusArrival struct {
	EstimatedArrival string
	Load             string
}

func (ba *Bus) Request() error {
	u := buildURL(ba)
	return ba.request(u)
}

func (ba *Bus) request(u string) (err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("AccountKey", dbcon.Conf.BusKey)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(ba)
	return nil
}

func buildURL(ba *Bus) string {
	u, _ := url.Parse("http://datamall2.mytransport.sg/ltaodataservice/BusArrival")
	u.Scheme = "http"
	v := make(url.Values)

	v.Add("BusStopID", ba.StationCode)
	v.Add("ServiceNo", ba.BusCode)

	u.RawQuery = v.Encode()
	return u.String()
}
