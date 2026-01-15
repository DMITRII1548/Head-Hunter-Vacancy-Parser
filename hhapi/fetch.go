package hhapi

import (
	"encoding/json"
	"fmt"
	"hhparser/config"
	"hhparser/entity"
	"net/http"
	"net/url"
	"time"
)

func getUrls() []string {
	var urls []string

	for i := 0; i < config.Page; i++ {
		params := url.Values{}
		params.Add("page", fmt.Sprint(i))
		params.Add("per_page", fmt.Sprint(config.PerPage))
		params.Add("area", fmt.Sprint(config.Area))

		url := config.BaseUrl + "?" + params.Encode()

		urls = append(urls, url)
	}

	return urls
}

func GetVacancies() (entity.VacancyResponse, error) {
	var allVacancies entity.VacancyResponse

	urls := getUrls()

	for _, url := range urls {
		// Rate limit
		time.Sleep(config.Delay * time.Millisecond)

		var vacancies entity.VacancyResponse

		res, err := http.Get(url)

		if err != nil {
			return entity.VacancyResponse{}, err
		}

		if res.StatusCode != http.StatusOK {
			return entity.VacancyResponse{}, fmt.Errorf("Filed response with status %s", res.Status)
		}

		err = json.NewDecoder(res.Body).Decode(&vacancies)

		res.Body.Close()

		if err != nil {
			fmt.Println(err)
			return entity.VacancyResponse{}, err
		}

		allVacancies.Items = append(allVacancies.Items, vacancies.Items...)
	}

	return allVacancies, nil
}