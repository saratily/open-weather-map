package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/kr/pretty"
)

type TemperatureDescription string

const (
	hot      TemperatureDescription = "HOT"
	cold     TemperatureDescription = "COLD"
	moderate TemperatureDescription = "MODERATE"
)

type Weather struct {
	Location  string `json:"name"`
	Condition []struct {
		Weather     string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Coordinate struct {
		Lat float32 `json:"lat"`
		Lon float32 `json:"lon"`
	} `json:"coord"`
	Temparature struct {
		Description    TemperatureDescription
		Temperature    float32 `json:"temp"`
		MinTemparature float32 `json:"temp_min"`
		MaxTemparature float32 `json:"temp_max"`
	} `json:"main"`
}

func NewWeather(config Config, lat, lon float64) (*Weather, error) {

	baseURL, err := url.Parse(config.Url)
	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{}
	params.Add("lat", fmt.Sprintf("%f", lat))
	params.Add("lon", fmt.Sprintf("%f", lon))
	params.Add("appid", config.ApiKey)
	baseURL.RawQuery = params.Encode()

	fmt.Println(baseURL)

	// Create the HTTP request
	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weather Weather
	err = json.Unmarshal([]byte(body), &weather)

	weather.setTemparature()

	if err != nil {
		log.Fatal(err)
	}

	pretty.Print(weather)
	return &weather, nil

}

func (w *Weather) setTemparature() {
	if w.Temparature.MaxTemparature > 305 {
		w.Temparature.Description = hot
	} else if w.Temparature.MinTemparature < 285 {
		w.Temparature.Description = cold
	} else {
		w.Temparature.Description = moderate
	}
}
