package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/timpark0807/go-tsla/handler"
)

func main() {
	// stations := helper.ReadFile()
	// helper.LoadDB(stations.Stations)

	router := mux.NewRouter()
	router.HandleFunc("/api/stations", handler.ListStations).Methods("GET")
	router.HandleFunc("/api/stations/zipcode={zipcode}", handler.GetStationByZip).Methods("GET")
	router.HandleFunc("/api/stations/state={state}", handler.GetStationByState).Methods("GET")

	// // router.HandleFunc("/api/station/{name}", handler.GetStation}).Methods("GET")
	router.HandleFunc("/api/stations/status={status}", handler.ListStatusStations).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
