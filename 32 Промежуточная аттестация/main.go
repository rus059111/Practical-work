/* Цель работы

Проверить и закрепить знания, полученные на курсе «Go-разработчик»:

    основы синтаксиса языка;
    условные операторы и циклы;
    работа с файловой системой;
    структуры данных;
    сериализация;
    многопоточность;
    обмен данными по сети.


Что нужно сделать

Вам нужно разработать сервис, предоставляющий информацию о городах. Данные хранятся в файле. В момент старта сервиса данные из файла кешируются в память, в момент завершения работы сервиса данные перезаписываются обратно в файл.


В каждой строке файла через запятую перечислена информация о городе:

    id (уникальный номер);
    name (название города);
    region (регион);
    district (округ);
    population (численность населения);
    foundation (год основания).


Требуется реализовать сервис имеющий следующий функционал:

    получение информации о городе по его id;
    добавление новой записи в список городов;
    удаление информации о городе по указанному id;
    обновление информации о численности населения города по указанному id;
    получение списка городов по указанному региону;
    получение списка городов по указанному округу;
    получения списка городов по указанному диапазону численности населения;
    получения списка городов по указанному диапазону года основания.


Что оценивается

    Безопасная работа с данными в памяти.
    Для завершения работы используется паттерн Graceful Shutdown.
    Заголовки HTTP-запросов соответствуют выполняемым операциям.
    Ответы содержат правильные коды HTTP-состояний.
    Для кодирования данных используется формат JSON.


Пример тестовых данных


Содержимое файла cities.csv

490,Москва,Москва,Центральный,11514330,1147

781,Санкт-Петербург,Санкт-Петербург,Северо-Западный,4848742,1703

634,Новосибирск,Новосибирская область,Сибирский,1498921,1893

829,Екатеринбург,Свердловская область,Уральский,1377738,1723

606,Нижний Новгород,Нижегородская область,Приволжский,1250615,1221

769,Самара,Самарская область,Приволжский,1164900,1586

643,Омск,Омская область,Сибирский,1154000,1716

922,Казань,Татарстан,Приволжский,1143546,1005

1058,Челябинск,Челябинская область,Уральский,1130273,1736

744,Ростов-на-Дону,Ростовская область,Южный,1091544,1749

62,Уфа,Башкортостан,Приволжский,1062300,1574

121,Волгоград,Волгоградская область,Южный,1021244,1589

693,Пермь,Пермский край,Приволжский,1000679,1723

410,Красноярск,Красноярский край,Сибирский,1000000,1628

159,Воронеж,Воронежская область,Центральный,889989,1586

797,Саратов,Саратовская область,Приволжский,836900,1590

380,Краснодар,Краснодарский край,Южный,744933,1793

771,Тольятти,Самарская область,Приволжский,719484,1737

993,Ижевск,Удмуртия,Приволжский,628117,1760

1002,Ульяновск,Ульяновская область,Приволжский,613793,1648

5,Барнаул,Алтайский край,Сибирский,612091,1730

704,Владивосток,Приморский край,Дальневосточный,592069,1860

1109,Ярославль,Ярославская область,Центральный,591486,1010

222,Иркутск,Иркутская область,Сибирский,587225,1661

989,Тюмень,Тюменская область,Уральский,581758,1586

177,Махачкала,Дагестан,Северо-Кавказский,577990,1844

1009,Хабаровск,Хабаровский край,Дальневосточный,577668,1858

653,Оренбург,Оренбургская область,Приволжский,570329,1743

321,Новокузнецк,Кемеровская область,Сибирский,547885,1618

315,Кемерово,Кемеровская область,Сибирский,532884,1918



Критерии оценки работы

Зачет — выполнены функциональные и технологические требования.

Незачёт  — выполнены только функциональные или технологические требования. */

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
