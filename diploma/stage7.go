package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func main() {
	data, err := getSupportData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Support Data:")
	for _, d := range data {
		fmt.Printf("Topic: %s, Active Tickets: %d\n", d.Topic, d.ActiveTickets)
	}
}

func getSupportData() ([]SupportData, error) {
	url := "http://127.0.0.1:8383/support"

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

	var supportData []SupportData
	err = json.Unmarshal(body, &supportData)
	if err != nil {
		return nil, err
	}

	return supportData, nil
}
