package model

type Stations struct {
	Stations []Station `json:"stations"`
}

type Station struct {
	Name    string  `json:"name"`
	URL     string  `json:"url"`
	Address Address `json:"address"`
	Active  bool    `json:"active"`
}

type Address struct {
	Street   string `json:"street"`
	Locality string `json:"locality"`
}
