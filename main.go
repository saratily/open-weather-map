package main

import (
	"fmt"
	"log"
	"open-weather-map/weather"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

func main() {

	if len(os.Args) != 3 {
		log.Fatal("Incorrect input format. Please use the format: go run .\\main.go <lag> <lon>")
	}

	lat, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil {
		// Handle the error, for example:
		fmt.Println("Error:", err)
		return
	}

	lon, err := strconv.ParseFloat(os.Args[2], 32)
	if err != nil {
		// Handle the error, for example:
		fmt.Println("Error:", err)
		return
	}

	// Read the YAML file
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the YAML data into a struct
	var config weather.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	weather.NewWeather(config, lat, lon)
}
