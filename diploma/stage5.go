package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

func main() {
	// Чтение файла
	filePath := "email.csv"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Разделение на строки
	lines := strings.Split(string(content), "\n")

	// Обработка строк
	var data []EmailData
	for _, line := range lines {
		fields := strings.Split(line, ";")

		// Проверка количества полей
		if len(fields) != 3 {
			continue
		}

		// Проверка кода страны и провайдера
		if !isValidCountry(fields[0]) || !isValidProvider(fields[1]) {
			continue
		}

		// Преобразование времени доставки в int
		deliveryTime, err := strconv.Atoi(fields[2])
		if err != nil {
			continue
		}

		// Создание структуры и добавление в результат
		item := EmailData{
			Country:      fields[0],
			Provider:     fields[1],
			DeliveryTime: deliveryTime,
		}
		data = append(data, item)
	}

	// Вывод результатов
	for _, item := range data {
		fmt.Printf("Country: %s, Provider: %s, DeliveryTime: %d\n",
			item.Country, item.Provider, item.DeliveryTime)
	}
}

func isValidCountry(country string) bool {
	validCountries := map[string]bool{
		"RU": true,
		"US": true,
		"DE": true,
		// Добавьте остальные коды стран
	}

	_, ok := validCountries[country]
	return ok
}

func isValidProvider(provider string) bool {
	validProviders := map[string]bool{
		"Gmail":       true,
		"Yahoo":       true,
		"Hotmail":     true,
		"MSN":         true,
		"Orange":      true,
		"Comcast":     true,
		"AOL":         true,
		"Live":        true,
		"RediffMail":  true,
		"GMX":         true,
		"Proton Mail": true,
		"Yandex":      true,
		"Mail.ru":     true,
		// Добавьте остальных провайдеров
	}

	_, ok := validProviders[provider]
	return ok
}
