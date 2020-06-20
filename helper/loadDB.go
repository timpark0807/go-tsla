package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/timpark0807/go-tsla/model"
)

func ReadFile() model.Stations {
	jsonFile, err := os.Open("data.json")
	var stations model.Stations

	if err != nil {
		fmt.Println(err)
		return stations
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &stations)

	defer jsonFile.Close()
	return stations
}

func LoadDB(stations []model.Station) {

	err := godotenv.Load()
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	valueStrings := make([]string, 0, len(stations))
	valueArgs := make([]interface{}, 0, len(stations)*8)
	for index, station := range stations {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?)")
		output := processLocality(station.Address.Locality)
		valueArgs = append(valueArgs, index)
		valueArgs = append(valueArgs, station.Name)
		valueArgs = append(valueArgs, station.URL)
		valueArgs = append(valueArgs, station.Active)
		valueArgs = append(valueArgs, station.Address.Street)
		valueArgs = append(valueArgs, output["city"])
		valueArgs = append(valueArgs, output["state"])
		valueArgs = append(valueArgs, output["zipcode"])
	}
	stmt := fmt.Sprintf("INSERT INTO superchargers (id, name, url, active, street, city, state, zipcode) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err = db.Exec(stmt, valueArgs...)
}

func processLocality(locality string) map[string]string {
	address := make(map[string]string)

	words := strings.Split(locality, ",")
	address["city"] = words[0]

	for _, word := range strings.Split(words[1], " ") {
		if strings.Contains(word, "-") {
			address["zipcode"] = strings.Split(word, "-")[0]
		} else if len(word) == 2 {
			address["state"] = word
		} else {
			address["zipcode"] = word
		}
	}
	return address
}
