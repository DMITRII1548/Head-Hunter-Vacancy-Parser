package entity

type Vacancy struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary *struct {
		Min *int `json:"from"`
		Max *int `json:"to"`
	} `json:"salary"`
	Url         string `json:"alternate_url"`
	Information struct {
		Requirement    *string `json:"requirement"`
		Responsibility *string `json:"responsibility"`
	} `json:"snippet"`
	Address *struct {
		City   *string `json:"city"`
		Street *string `json:"street"`
		FullAddress *string `json:"raw"`
		Latitude *float64 `json:"lat"`
		Longitude *float64 `json:"lng"`
	} `json:"address"`
}