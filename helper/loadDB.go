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
	valueArgs := make([]interface{}, 0, len(stations)*6)
	for index, station := range stations {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, index)
		valueArgs = append(valueArgs, station.Name)
		valueArgs = append(valueArgs, station.URL)
		valueArgs = append(valueArgs, station.Active)
		valueArgs = append(valueArgs, station.Address.Street)
		valueArgs = append(valueArgs, station.Address.Locality)
	}
	stmt := fmt.Sprintf("INSERT INTO superchargers (id, name, url, active, street, locality) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err = db.Exec(stmt, valueArgs...)
	fmt.Println(err)

}
