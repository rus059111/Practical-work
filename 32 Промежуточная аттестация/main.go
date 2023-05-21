package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type City struct {
	ID         int
	Name       string
	Region     string
	District   string
	Population int
	Foundation int
}

var cityData []City

func main() {
	// Загрузка данных из файла
	err := loadDataFromFile("cities.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Регистрация обработчиков запросов
	http.HandleFunc("/city", getCityByIDHandler)
	http.HandleFunc("/add", addCityHandler)
	http.HandleFunc("/delete", deleteCityByIDHandler)
	http.HandleFunc("/update/population", updatePopulationByIDHandler)
	http.HandleFunc("/cities/region", getCitiesByRegionHandler)
	http.HandleFunc("/cities/district", getCitiesByDistrictHandler)
	http.HandleFunc("/cities/population-range", getCitiesByPopulationRangeHandler)
	http.HandleFunc("/cities/foundation-range", getCitiesByFoundationRangeHandler)

	// Запуск HTTP-сервера
	log.Println("Сервис городов запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// Запись данных обратно в файл при завершении работы сервиса
	err = saveDataToFile("cities.csv")
	if err != nil {
		log.Fatal(err)
	}
}

func loadDataFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.Comment = '#'

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	cityData = make([]City, len(records))
	for i, record := range records {
		id, _ := strconv.Atoi(record[0])
		population, _ := strconv.Atoi(record[4])
		foundation, _ := strconv.Atoi(record[5])

		cityData[i] = City{
			ID:         id,
			Name:       record[1],
			Region:     record[2],
			District:   record[3],
			Population: population,
			Foundation: foundation,
		}
	}

	return nil
}

func saveDataToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ','

	for _, city := range cityData {
		record := []string{
			strconv.Itoa(city.ID),
			city.Name,
			city.Region,
			city.District,
			strconv.Itoa(city.Population),
			strconv.Itoa(city.Foundation),
		}
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}

func getCityByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid city ID", http.StatusBadRequest)
		return
	}

	for _, city := range cityData {
		if city.ID == id {
			fmt.Fprintf(w, "ID: %d, Name: %s, Region: %s, District: %s, Population: %d, Foundation: %d\n",
				city.ID, city.Name, city.Region, city.District, city.Population, city.Foundation)
			return
		}
	}

	http.NotFound(w, r)
}

func addCityHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	region := r.URL.Query().Get("region")
	district := r.URL.Query().Get("district")
	population, err := strconv.Atoi(r.URL.Query().Get("population"))
	if err != nil {
		http.Error(w, "Invalid population", http.StatusBadRequest)
		return
	}
	foundation, err := strconv.Atoi(r.URL.Query().Get("foundation"))
	if err != nil {
		http.Error(w, "Invalid foundation year", http.StatusBadRequest)
		return
	}

	id := len(cityData) + 1
	city := City{
		ID:         id,
		Name:       name,
		Region:     region,
		District:   district,
		Population: population,
		Foundation: foundation,
	}
	cityData = append(cityData, city)

	fmt.Fprintf(w, "City added with ID: %d\n", id)
}

func deleteCityByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid city ID", http.StatusBadRequest)
		return
	}

	for i, city := range cityData {
		if city.ID == id {
			cityData = append(cityData[:i], cityData[i+1:]...)
			fmt.Fprintf(w, "City deleted with ID: %d\n", id)
			return
		}
	}

	http.NotFound(w, r)
}

func updatePopulationByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid city ID", http.StatusBadRequest)
		return
	}

	population, err := strconv.Atoi(r.URL.Query().Get("population"))
	if err != nil {
		http.Error(w, "Invalid population", http.StatusBadRequest)
		return
	}

	for i, city := range cityData {
		if city.ID == id {
			cityData[i].Population = population
			fmt.Fprintf(w, "Population updated for city with ID: %d\n", id)
			return
		}
	}

	http.NotFound(w, r)
}

func getCitiesByRegionHandler(w http.ResponseWriter, r *http.Request) {
	region := r.URL.Query().Get("region")

	for _, city := range cityData {
		if strings.EqualFold(city.Region, region) {
			fmt.Fprintf(w, "ID: %d, Name: %s, Region: %s, District: %s, Population: %d, Foundation: %d\n",
				city.ID, city.Name, city.Region, city.District, city.Population, city.Foundation)
		}
	}
}

func getCitiesByDistrictHandler(w http.ResponseWriter, r *http.Request) {
	district := r.URL.Query().Get("district")

	for _, city := range cityData {
		if strings.EqualFold(city.District, district) {
			fmt.Fprintf(w, "ID: %d, Name: %s, Region: %s, District: %s, Population: %d, Foundation: %d\n",
				city.ID, city.Name, city.Region, city.District, city.Population, city.Foundation)
		}
	}
}

func getCitiesByPopulationRangeHandler(w http.ResponseWriter, r *http.Request) {
	minPopulation, err := strconv.Atoi(r.URL.Query().Get("min_population"))
	if err != nil {
		http.Error(w, "Invalid minimum population", http.StatusBadRequest)
		return
	}

	maxPopulation, err := strconv.Atoi(r.URL.Query().Get("max_population"))
	if err != nil {
		http.Error(w, "Invalid maximum population", http.StatusBadRequest)
		return
	}

	for _, city := range cityData {
		if city.Population >= minPopulation && city.Population <= maxPopulation {
			fmt.Fprintf(w, "ID: %d, Name: %s, Region: %s, District: %s, Population: %d, Foundation: %d\n",
				city.ID, city.Name, city.Region, city.District, city.Population, city.Foundation)
		}
	}
}

func getCitiesByFoundationRangeHandler(w http.ResponseWriter, r *http.Request) {
	minFoundation, err := strconv.Atoi(r.URL.Query().Get("min_foundation"))
	if err != nil {
		http.Error(w, "Invalid minimum foundation year", http.StatusBadRequest)
		return
	}

	maxFoundation, err := strconv.Atoi(r.URL.Query().Get("max_foundation"))
	if err != nil {
		http.Error(w, "Invalid maximum foundation year", http.StatusBadRequest)
		return
	}

	for _, city := range cityData {
		if city.Foundation >= minFoundation && city.Foundation <= maxFoundation {
			fmt.Fprintf(w, "ID: %d, Name: %s, Region: %s, District: %s, Population: %d, Foundation: %d\n",
				city.ID, city.Name, city.Region, city.District, city.Population, city.Foundation)
		}
	}
}
