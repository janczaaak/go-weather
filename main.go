package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC    float64 `json:"temp_c"`
		TempF    float64 `json:"temp_f"`
		WindMph  float64 `json:"wind_mph"`
		WindKph  float64 `json:"wind_kph"`
		PrecipMm float64 `json:"precip_mm"`
		PrecipIn float64 `json:"precip_in"`
		Cloud    int     `json:"cloud"`
	} `json:"current"`
}

func main() {
	var city string
	fmt.Println("Hey, which city you want me to check: ")
	fmt.Scanf("%s", &city)

	weather, err := getWeather(city)
	if err != nil {
		fmt.Println("Couldnt download data")
		return
	}

	printWeather(weather)
}

func getWeather(city string) (*Weather, error) {
	apikey := "weather_api_key"
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apikey, city)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}

func printWeather(weather *Weather) {
	fmt.Println("------------------")
	fmt.Printf("City: %s\nCountry: %s\n", weather.Location.Name, weather.Location.Country)
	fmt.Println("------------------")
	fmt.Printf("Temp C: %.2f\nTemp F: %.2f\n", weather.Current.TempC, weather.Current.TempF)
	fmt.Printf("Precipitation Mm: %.2f\nPercipitation In: %.2f\n", weather.Current.PrecipMm, weather.Current.PrecipIn)
	fmt.Printf("Wind Kph: %.2f\nWind Mph: %.2f\n", weather.Current.WindKph, weather.Current.WindMph)
	fmt.Printf("Cloud: %v\n", weather.Current.Cloud)
	fmt.Println("------------------")
}
