package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type CSVData struct {
	Country      string
	Load         int
	ResponseTime int
	Provider     string
	Stability    float32
	TTFB         int
	CallDuration int
}

func main() {
	// Чтение файла
	filePath := "VoiceCall.csv"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Разделение на строки
	lines := strings.Split(string(content), "\n")

	// Обработка строк
	var data []CSVData
	for _, line := range lines {
		fields := strings.Split(line, ";")

		// Проверка количества полей
		if len(fields) != 8 {
			continue
		}

		// Разбор полей
		load, err := strconv.Atoi(fields[1])
		if err != nil {
			continue
		}

		responseTime, err := strconv.Atoi(fields[2])
		if err != nil {
			continue
		}

		stability, err := strconv.ParseFloat(fields[4], 32)
		if err != nil {
			continue
		}

		ttfb, err := strconv.Atoi(fields[5])
		if err != nil {
			continue
		}

		callDuration, err := strconv.Atoi(fields[6])
		if err != nil {
			continue
		}

		// Проверка кода страны и провайдера
		if !isValidCountry(fields[0]) || !isValidProvider(fields[3]) {
			continue
		}

		// Создание структуры и добавление в результат
		item := CSVData{
			Country:      fields[0],
			Load:         load,
			ResponseTime: responseTime,
			Provider:     fields[3],
			Stability:    float32(stability),
			TTFB:         ttfb,
			CallDuration: callDuration,
		}
		data = append(data, item)
	}

	// Вывод результатов
	for _, item := range data {
		fmt.Printf("Country: %s, Load: %d, ResponseTime: %d, Provider: %s, Stability: %.2f, TTFB: %d, CallDuration: %d\n",
			item.Country, item.Load, item.ResponseTime, item.Provider, item.Stability, item.TTFB, item.CallDuration)
	}
}

func isValidCountry(country string) bool {
	validCountries := map[string]bool{
		"MC": true,
		"US": true,
		"DE": true,
		// Добавьте остальные коды стран
	}

	_, ok := validCountries[country]
	return ok
}

func isValidProvider(provider string) bool {
	validProviders := map[string]bool{
		"TransparentCalls": true,
		"E-Voice":          true,
		"JustPhone":        true,
		// Добавьте остальных провайдеров
	}

	_, ok := validProviders[provider]
	return ok
}
