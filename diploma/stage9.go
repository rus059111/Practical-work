package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

func main() {
	// Создание роутера с использованием gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/", HandleConnection).Methods("GET")

	// Создание HTTP-сервера
	server := &http.Server{
		Addr:    "localhost:8282",
		Handler: router,
	}

	// Запуск сервера
	log.Println("Сервер запущен на http://localhost:8282")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	// Отправка GET-запроса
	url := "http://127.0.0.1:8585"
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "API request failed", http.StatusInternalServerError)
		return
	}

	// Чтение тела ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Разбор JSON в структуры IncidentData
	var incidentData []IncidentData
	err = json.Unmarshal(body, &incidentData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
