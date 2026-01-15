package helper

import (
	"fmt"
	"hhparser/entity"
)

func StringMinSalary(vacancy entity.Vacancy) string {
	if vacancy.Salary == nil || vacancy.Salary.Min == nil {
		return ""
	}

	return fmt.Sprint(vacancy.Salary.Min)
}

func StringMaxSalary(vacancy entity.Vacancy) string {
	if vacancy.Salary == nil || vacancy.Salary.Max == nil {
		return ""
	}

	return fmt.Sprint(vacancy.Salary.Max)
}

func StringRequirementInformation(vacancy entity.Vacancy) string {
	if vacancy.Information.Requirement == nil {
		return ""
	}

	return *vacancy.Information.Requirement
}

func StringResponsibilityInformation(vacancy entity.Vacancy) string {
	if vacancy.Information.Responsibility == nil {
		return ""
	}

	return *vacancy.Information.Responsibility
}

func StringCity(vacancy entity.Vacancy) string {
	if vacancy.Address == nil || vacancy.Address.City == nil {
		return ""
	}

	return *vacancy.Address.City
}

func StringStreet(vacancy entity.Vacancy) string {
	if vacancy.Address == nil || vacancy.Address.Street == nil {
		return ""
	}

	return *vacancy.Address.Street
}

func StringFullAddress(vacancy entity.Vacancy) string {
	if vacancy.Address == nil || vacancy.Address.FullAddress == nil {
		return ""
	}

	return *vacancy.Address.FullAddress
}

func StringLatitude(vacancy entity.Vacancy) string {
	if vacancy.Address == nil || vacancy.Address.Latitude == nil {
		return ""
	}

	return fmt.Sprint(*vacancy.Address.Latitude)
}

func StringLongitude(vacancy entity.Vacancy) string {
	if vacancy.Address == nil || vacancy.Address.Longitude == nil {
		return ""
	}

	return fmt.Sprint(*vacancy.Address.Longitude)
}