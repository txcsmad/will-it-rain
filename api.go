package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Data struct {
	PrecipitationProbability float64 `json:"precipProbability"`
}

type Daily struct {
	Data []Data `json:"data"`
}

type Forecast struct {
	Daily Daily `json:"daily"`
}

func getForecast() (*Forecast, error) {
	res, err := http.Get(fmt.Sprintf("https://api.forecast.io/forecast/%s/%f,%f", forecastIOKey, austinLatitude, austinLongitude))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("get forecast: bad status code")
	}

	var f Forecast
	if err := json.NewDecoder(res.Body).Decode(&f); err != nil {
		return nil, err
	}
	return &f, nil
}
