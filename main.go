package main

import "github.com/carlqt/ez-bus/geocode"

type GeoRevResponse struct {
}

func main() {
	geocode.Reverse(1.123, 1.32421)
}
