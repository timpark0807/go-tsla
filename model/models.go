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
	Locality string `json:"locality,omitempty"`
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zipcode  string `json:"zipcode"`
}
