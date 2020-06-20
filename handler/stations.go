package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/timpark0807/go-tsla/model"
)

func ListStations(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error")
	}

	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	db, _ := sql.Open("mysql", dbConnectionString)
	defer db.Close()

	results, err := db.Query("SELECT name, url, active, street, city, state, zipcode FROM superchargers")

	var output []model.Station
	for results.Next() {
		var station model.Station
		err := results.Scan(&station.Name, &station.URL, &station.Active, &station.Address.Street, &station.Address.City, &station.Address.State, &station.Address.Zipcode)
		if err == nil {
			output = append(output, station)
		}
	}
	json.NewEncoder(w).Encode(output)
}

func GetStation(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error")
	}
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	db, _ := sql.Open("mysql", dbConnectionString)
	defer db.Close()

	var params = mux.Vars(r)
	var station model.Station
	db.QueryRow("SELECT name, url, active, street, city, state, zipcode FROM superchargers where zipcode = ?", params["zipcode"]).Scan(&station.Name, &station.URL, &station.Active, &station.Address.Street, &station.Address.City, &station.Address.State, &station.Address.Zipcode)
	json.NewEncoder(w).Encode(station)

}

func ListStatusStations(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error")
	}

	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	db, _ := sql.Open("mysql", dbConnectionString)
	defer db.Close()

	var params = mux.Vars(r)
	results, err := db.Query("SELECT name, url, active, street, city, state, zipcode FROM superchargers where active=?", params["active"])

	var output []model.Station
	for results.Next() {
		var station model.Station
		err := results.Scan(&station.Name, &station.URL, &station.Active, &station.Address.Street, &station.Address.City, &station.Address.State, &station.Address.Zipcode)
		if err == nil {
			output = append(output, station)
		}
	}
	json.NewEncoder(w).Encode(output)
}
