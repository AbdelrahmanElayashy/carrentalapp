package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentalmanagement/logic/model"
	"time"

	log "github.com/sirupsen/logrus"
)

type CarAPI struct {
	apiURL       string
	CarsEndpoint string
}

type GetCarsApiResponse struct {
	Cars []model.Car `json:"cars"`
}

func NewCarAPI(apiURL string) *CarAPI {
	return &CarAPI{
		apiURL:       apiURL,
		CarsEndpoint: "/cars",
	}
}

func (c *CarAPI) ListAllCars() ([]model.Car, error) {
	var msg string
	apiURL := fmt.Sprintf("%s%s", c.apiURL, c.CarsEndpoint)

	// Make the GET request to the API
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(apiURL)
	if err != nil {
		msg = "Failed to make GET request to the Car API"
		log.Error(msg, ": ", err)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		msg = fmt.Sprintf("API request failed with status code: %d", resp.StatusCode)
		log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	// Decode the JSON response
	var getCarsApiResponse GetCarsApiResponse
	err = json.NewDecoder(resp.Body).Decode(&getCarsApiResponse)
	if err != nil {
		msg = "Failed to decode JSON response from Car API"
		log.Error(msg, ": ", err)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	return getCarsApiResponse.Cars, nil
}

func (c *CarAPI) CarExists(vin model.Vin) (bool, error) {
	var msg string
	apiURL := fmt.Sprintf("%s%s/%s", c.apiURL, c.CarsEndpoint, vin.Vin)

	// Make the GET request to the API
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(apiURL)
	if err != nil {
		msg = fmt.Sprintf("Failed to make GET request for car with VIN %s", vin)
		log.Error(msg, ": ", err)
		return false, fmt.Errorf("%s: %w", msg, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg = fmt.Sprintf("API request for car with VIN %s failed with status code: %d", vin, resp.StatusCode)
		log.Warn(msg)
		return false, fmt.Errorf(msg)
	}

	return true, nil
}
