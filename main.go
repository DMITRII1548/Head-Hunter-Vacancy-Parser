package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

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

type VacancyResponse struct {
	Items []Vacancy `json:"items"`
}

const perPage = 100
const page = 20
const area = 113  // Russia
const delay = 200 // Milliseconds
const baseUrl = "https://api.hh.ru/vacancies"

func getUrls() []string {
	var urls []string

	for i := 0; i < page; i++ {
		params := url.Values{}
		params.Add("page", fmt.Sprint(i))
		params.Add("per_page", fmt.Sprint(perPage))
		params.Add("area", fmt.Sprint(area))

		url := baseUrl + "?" + params.Encode()

		urls = append(urls, url)
	}

	return urls
}

func getVacancies(urls []string) (VacancyResponse, error) {
	var allVacancies VacancyResponse
	
	for _, url := range urls {
		// Rate limit
		fmt.Println(url)
		time.Sleep(delay * time.Millisecond)

		var vacancies VacancyResponse

		res, err := http.Get(url)

		if err != nil {
			return VacancyResponse{}, err
		}

		if res.StatusCode != http.StatusOK {
			return VacancyResponse{}, fmt.Errorf("Filed response with status ", res.Status)
		}

		err = json.NewDecoder(res.Body).Decode(&vacancies)

		res.Body.Close()

		if err != nil {
			fmt.Println(err)
			return VacancyResponse{}, err
		}

		allVacancies.Items = append(allVacancies.Items, vacancies.Items...)
	}

	return allVacancies, nil
}

func getStringMinSalary(vacancy Vacancy) string {
	if vacancy.Salary == nil || vacancy.Salary.Min == nil {
		return ""
	}

	return fmt.Sprint(vacancy.Salary.Min)
}

func getStringMaxSalary(vacancy Vacancy) string {
	if vacancy.Salary == nil || vacancy.Salary.Max == nil {
		return ""
	}

	return fmt.Sprint(vacancy.Salary.Max)
}

func getStringRequirementInformation(vacancy Vacancy) string {
	if vacancy.Information.Requirement == nil {
		return ""
	}
	
	return *vacancy.Information.Requirement
}

func getStringResponsibilityInformation(vacancy Vacancy) string {
	if vacancy.Information.Responsibility == nil {
		return ""
	}

	return *vacancy.Information.Responsibility
}

func getStringCity(vacancy Vacancy) string {
	if vacancy.Address == nil || vacancy.Address.City == nil {
		return ""
	}

	return *vacancy.Address.City
}

func getStringStreet(vacancy Vacancy) string {
	if vacancy.Address == nil || vacancy.Address.Street == nil {
		return ""
	}

	return *vacancy.Address.Street
}

func getStringFullAddress(vacancy Vacancy) string {
	if vacancy.Address == nil || vacancy.Address.FullAddress == nil {
		return ""
	}

	return *vacancy.Address.FullAddress
}

func getStringLatitude(vacancy Vacancy) string {
	if vacancy.Address == nil || vacancy.Address.Latitude == nil {
		return ""
	}

	return fmt.Sprint(*vacancy.Address.Latitude)
}

func getStringLongitude(vacancy Vacancy) string {
	if vacancy.Address == nil || vacancy.Address.Longitude == nil {
		return ""
	}

	return fmt.Sprint(*vacancy.Address.Longitude)
}

func saveToCSV(filename string, vacancies []Vacancy) error {
	filePath := "./storage/" + filename + ".csv"
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	csvHeaders := []string{
		"id",
		"name",
		"salary_min",
		"salary_max",
		"url",
		"requirement",
		"responsibility",
		"city",
		"street",
		"address",
		"latitude",
		"longitude",
	}

	writer.Write(csvHeaders)

	for _, vacancy := range vacancies {
		record := []string{
			vacancy.ID,
			vacancy.Name,
			getStringMinSalary(vacancy),
			getStringMaxSalary(vacancy),
			vacancy.Url,
			getStringRequirementInformation(vacancy),
			getStringResponsibilityInformation(vacancy),
			getStringCity(vacancy),
			getStringStreet(vacancy),
			getStringFullAddress(vacancy),
			getStringLatitude(vacancy),
			getStringLongitude(vacancy),
		}

		writer.Write(record)
	}

	fmt.Println("CVS created successful: ", filePath)

	return nil
}

func main() {
	urls := getUrls()
	vacancies, err := getVacancies(urls)

	if err != nil {
		fmt.Println(err)
	}

	err = saveToCSV(fmt.Sprint(time.Now().UnixNano()), vacancies.Items)

	if err != nil {
		fmt.Println(err)
	}
}
