package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/carlqt/ez-bus/env"
)

type Bus struct {
	BusCode     string
	StationCode string `json:"BusStopID"`
	Services    []*Service
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
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("AccountKey", env.Conf.BusKey)
	req.Header.Set("Accept", "application/json")

	resp, err := env.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(ba)

	for _, service := range ba.Services {
		service.formatTimeDuration()
	}
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

func (s *Service) formatTimeDuration() {
	format := time.RFC3339
	parsedTime, _ := time.Parse(format, s.NextBus.EstimatedArrival)
	duration := time.Until(parsedTime)
	s.NextBus.EstimatedArrival = fmt.Sprintf("%f", duration.Minutes())

	parsedTime, _ = time.Parse(format, s.SubsequentBus.EstimatedArrival)
	duration = time.Until(parsedTime)
	s.SubsequentBus.EstimatedArrival = fmt.Sprintf("%f", duration.Minutes())
}
