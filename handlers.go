package main

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func NearbyStations(w http.ResponseWriter, r *http.Request) {
	// q := r.URL.Query()
}
