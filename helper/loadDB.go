package helper

import (
	"fmt"

	"github.com/timpark0807/go-tsla/model"
)

func LoadDB(stations model.Stations) {

	for i := 0; i < len(stations.Stations); i++ {
		fmt.Println(stations.Stations[i])
	}
}
