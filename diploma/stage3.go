package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func main() {
	// Адрес для GET-запроса
	url := "http://127.0.0.1:8383/mms"

	// Отправка GET-запроса
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Проверка кода ответа
	if response.StatusCode != http.StatusOK {
		log.Fatalf("GET request failed with status: %s", response.Status)
	}

	// Чтение ответа в срез байтов
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Преобразование среза байтов в срез структур MMSData
	var data []MMSData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Фильтрация некорректных элементов
	data = filterData(data)

	// Вывод результатов
	for _, item := range data {
		fmt.Printf("Country: %s, Provider: %s, Bandwidth: %s, ResponseTime: %s\n",
			item.Country, item.Provider, item.Bandwidth, item.ResponseTime)
	}
}

func filterData(data []MMSData) []MMSData {
	var filteredData []MMSData

	for _, item := range data {
		if isValidCountry(item.Country) && isValidProvider(item.Provider) {
			filteredData = append(filteredData, item)
		}
	}

	return filteredData
}

func isValidCountry(country string) bool {
	validCountries := map[string]bool{
		"BG": true,
		"US": true,
		"DE": true,
		// Добавьте остальные страны
	}

	_, ok := validCountries[country]
	return ok
}

func isValidProvider(provider string) bool {
	validProviders := map[string]bool{
		"Topolo": true,
		"Rond":   true,
		"Kildy":  true,
		// Добавьте остальных провайдеров
	}

	_, ok := validProviders[provider]
	return ok
}
