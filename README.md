
# Open Weather Map

## Introduction

An http server in Go that uses the Open Weather API that exposes an endpoint that takes in lat/long coordinates. This endpoint should return what the weather condition is outside in that area (snow, rain, etc) and whether the temperature is hot, cold, or moderate (use your own discretion on what temperature range equates to each type).
 
The Open Weather API can be found here: https://openweathermap.org/api. 
 
Even though most of the API calls found on Open Weather API aren’t free, you should be able to use the free “current weather data”, https://openweathermap.org/current, API call for this project. First, sign-up for an account, which shouldn’t require credit card or
payment information. Once you’ve created an account, use https://openweathermap.org/faq to get your API Key setup to start using the API.


### Configure the project

- Configure your OpenWeatherMap API key in config.yaml
- If a import package is missing and throws, then run go mod tidy to fetch required packages.

## Run the project

To run the project, type the following command in cmd:

go run .\main.go <lat> <lon>

This program will take 2 arguments latitude and longitude. If these parameters are missing then this program will throw an error
