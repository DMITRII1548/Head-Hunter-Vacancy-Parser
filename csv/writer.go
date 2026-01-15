package csvwriter

import (
	"encoding/csv"
	"fmt"
	"hhparser/entity"
	"hhparser/helper"
	"os"
)

func SaveVacanciesToCSV(filename string, vacancies []entity.Vacancy) error {
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
			helper.StringMinSalary(vacancy),
			helper.StringMaxSalary(vacancy),
			vacancy.Url,
			helper.StringRequirementInformation(vacancy),
			helper.StringResponsibilityInformation(vacancy),
			helper.StringCity(vacancy),
			helper.StringStreet(vacancy),
			helper.StringFullAddress(vacancy),
			helper.StringLatitude(vacancy),
			helper.StringLongitude(vacancy),
		}

		writer.Write(record)
	}

	fmt.Println("CVS created successful: ", filePath)

	return nil
}
