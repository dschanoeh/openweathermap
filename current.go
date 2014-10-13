package openweathermap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// NewCurrent returns a new WeatherData pointer with the supplied
// arguments.
func NewCurrent(unit string) (*CurrentWeatherData, error) {
	unitChoice := strings.ToLower(unit)
	for _, i := range dataUnits {
		if strings.Contains(unitChoice, i) {
			return &CurrentWeatherData{Units: unitChoice}, nil
		}
	}
	return nil, errors.New("ERROR: unit of measure not available")
}

// CurrentByName will provide the current weather with the
// provided location name.
func (w *CurrentWeatherData) CurrentByName(location string) {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(baseUrl, "q=%s&units=%s"), location, w.Units))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(result, &w)
	if err != nil {
		log.Fatalln(err)
	}
}

// CurrentByCoordinates will provide the current weather with the
// provided location coordinates.
func (w *CurrentWeatherData) CurrentByCoordinates(location *Coordinates) {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(baseUrl, "lat=%f&lon=%f&units=%s"), location.Latitude, location.Longitude, w.Units))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(result, &w)
	if err != nil {
		log.Fatalln(err)
	}
}

// CurrentByID will provide the current weather with the
// provided location ID.
func (w *CurrentWeatherData) CurrentByID(id int) {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(baseUrl, "id=%d&units=%s"), id, w.Units))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(result, &w)
	if err != nil {
		log.Fatalln(err)
	}
}

func (w *CurrentWeatherData) GetByArea() {}
