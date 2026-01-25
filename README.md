# Head Hunter parser

## Description

This project parses up to 2000 vacancies from the Russia region and saves the results in CSV files.

It fetches data from the HH.ru API, converts nullable fields safely, and exports structured data.

## Project structure
HeadHunterParser/
├─ main.go        - Entry point of the application; orchestrates the workflow
├─ hhapi/         - Functions for interacting with the HH.ru API
├─ helpers/       - Utility functions for data conversion and nullable field handling
├─ entity/        - Definitions of data structures
├─ csv/           - Functions for exporting data to CSV files
├─ config/        - Project configuration and constants
└─ storage/       - Directory for storing generated CSV files


## Requirements
1) Golang >= 1.24.1

## Installation

1) git clone https://github.com/DMITRII1548/Head-Hunter-Vacancy-Parser.git
2) cd Head-Hunter-Vacancy-Parser
3) go mod download
4) mkdir storage
5) Run project - go run main.go 
6) Create a build - go build -o HeadHunterParser.exe main.go