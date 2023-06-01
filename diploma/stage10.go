package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}

type SMSData struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	Receiver string `json:"receiver"`
}

type MMSData struct {
	ID     int      `json:"id"`
	Images []string `json:"images"`
}

type VoiceCallData struct {
	ID     int    `json:"id"`
	Number string `json:"number"`
}

type EmailData struct {
	ID      int      `json:"id"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
	To      []string `json:"to"`
}

type BillingData struct {
	Total   float64 `json:"total"`
	Invoice string  `json:"invoice"`
}

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

func main() {
	result, err := getDataFromAPI()
	if err != nil {
		fmt.Println("Ошибка при получении данных:", err)
		return
	}

	fmt.Println("Статус:", result.Status)
	if result.Status {
		fmt.Println("Данные:")
		fmt.Println("SMS:", result.Data.SMS)
		fmt.Println("MMS:", result.Data.MMS)
		fmt.Println("Voice Call:", result.Data.VoiceCall)
		fmt.Println("Email:", result.Data.Email)
		fmt.Println("Billing:", result.Data.Billing)
		fmt.Println("Support:", result.Data.Support)
		fmt.Println("Incidents:", result.Data.Incidents)
	} else {
		fmt.Println("Ошибка:", result.Error)
	}
}

func getDataFromAPI() (ResultT, error) {
	url := "http://127.0.0.1:8585"

	resp, err := http.Get(url)
	if err != nil {
		return ResultT{}, fmt.Errorf("Ошибка при отправке GET-запроса: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ResultT{}, fmt.Errorf("Ошибка при чтении тела ответа: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return ResultT{}, fmt.Errorf("Ошибка при выполнении запроса: код состояния %v", resp.StatusCode)
	}

	var result ResultT
	err = json.Unmarshal(body, &result)
	if err != nil {
		return ResultT{}, fmt.Errorf("Ошибка при разборе JSON: %v", err)
	}

	return result, nil
}

func transformFieldName(fieldName string) string {
	words := strings.Split(fieldName, "")
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, "_")
}
