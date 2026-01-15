package main

import (
	"fmt"
	csvwriter "hhparser/csv"
	"hhparser/hhapi"
	"time"
)

func main() {
	vacancies, err := hhapi.GetVacancies()

	if err != nil {
		fmt.Println(err)
	}

	err = csvwriter.SaveVacanciesToCSV(fmt.Sprint(time.Now().UnixNano()), vacancies.Items)

	if err != nil {
		fmt.Println(err)
	}
}
