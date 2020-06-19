package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/timpark0807/go-tsla/helper"
	"github.com/timpark0807/go-tsla/model"
)

func readFile() model.Stations {
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

func main() {
	stations := readFile()
	helper.LoadDB(stations)
}
