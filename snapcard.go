package snapcard

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://api.snapcard.io"

func basicRequest(path string) ([]byte, error) {
	if path[0] != '/' {
		path = "/" + path
	}

	resp, err := http.Get(baseURL + path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

// Rate takes a trading pair and returns the rate as a float64
func Rate(pair string) (float64, error) {

	body, err := basicRequest("rates")
	if err != nil {
		return 0, err
	}

	rates := map[string]float64{}

	err = json.Unmarshal(body, &rates)
	if err != nil {
		return 0, err
	}

	if rate, ok := rates[pair]; !ok {
		return 0, errors.New("Invalid pair")
	} else {
		return rate, nil
	}

}
