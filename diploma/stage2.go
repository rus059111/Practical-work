package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func main() {
	// Путь к файлу CSV
	filePath := "SMS.csv"

	// Чтение файла
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Создание CSV reader
	reader := csv.NewReader(file)

	// Установка разделителя
	reader.Comma = ';'

	// Срез для хранения данных
	var data []SMSData

	// Чтение и разбор строк
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// Проверка количества полей
		if len(line) != 4 {
			continue
		}

		// Разбор строк на показатели
		country := line[0]
		bandwidth := line[1]
		responseTime := line[2]
		provider := line[3]

		// Проверка существования страны и провайдера
		if !isValidCountry(country) || !isValidProvider(provider) {
			continue
		}

		// Создание структуры SMSData
		smsData := SMSData{
			Country:      country,
			Bandwidth:    bandwidth,
			ResponseTime: responseTime,
			Provider:     provider,
		}

		// Добавление структуры в срез
		data = append(data, smsData)
	}

	// Вывод результатов
	for _, item := range data {
		fmt.Printf("Country: %s, Bandwidth: %s, ResponseTime: %s, Provider: %s\n",
			item.Country, item.Bandwidth, item.ResponseTime, item.Provider)
	}
}

func isValidCountry(country string) bool {
	// Мапа со списком корректных стран
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
	// Мапа со списком корректных провайдеров
	validProviders := map[string]bool{
		"Topolo": true,
		"Rond":   true,
		"Kildy":  true,
		// Добавьте остальных провайдеров
	}

	_, ok := validProviders[provider]
	return ok
}
