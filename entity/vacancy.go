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
		City        *string  `json:"city"`
		Street      *string  `json:"street"`
		FullAddress *string  `json:"raw"`
		Latitude    *float64 `json:"lat"`
		Longitude   *float64 `json:"lng"`
	} `json:"address"`
	Experience struct {
		Value string `json:"id"`
	} `json:"experience"`
	EmploymentForm struct {
		Value string `json:"id"`
	} `json:"employment_form"`
	Schedule struct {
		Value string `json:"id"`
	} `json:"schedule"`
	Format []struct {
		Value string `json:"id"`
	} `json:"work_format"`
	Profession []struct {
		Value string `json:"name"`
	} `json:"professional_roles"`
	Employer struct {
		LogoUrls *struct {
			Original string `json:"original"`
		} `json:"logo_urls"`
	} `json:"employer"`
}
