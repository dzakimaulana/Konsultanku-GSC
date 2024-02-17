package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func HitAPI(method string, url string, data map[string]interface{}) (map[string]interface{}, error) {

	// Encoding JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Create Request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// Create Client
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle the response here
	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return nil, err
	}

	// Check Status Code
	if resp.StatusCode == 200 {
		return responseBody, nil
	}
	errOutput := ErrorOutput(responseBody)
	return nil, errors.New(errOutput)
}
