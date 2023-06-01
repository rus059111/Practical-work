package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

func main() {
	data, err := getIncidentData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Incident Data:")
	for _, d := range data {
		fmt.Printf("Topic: %s, Status: %s\n", d.Topic, d.Status)
	}
}

func getIncidentData() ([]IncidentData, error) {
	url := "http://127.0.0.1:8383/accendent"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var incidentData []IncidentData
	err = json.Unmarshal(body, &incidentData)
	if err != nil {
		return nil, err
	}

	return incidentData, nil
}
