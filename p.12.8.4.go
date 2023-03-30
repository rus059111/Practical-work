package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	// Открываем файл для записи
	fileName := "log.txt"
	lineNumber := 1

	for {
		fmt.Print("Введите строку: ")
		var text string
		fmt.Scanln(&text)
		now := time.Now()
		if text == "exit" {
			break
		}

		data := []byte(fmt.Sprintf("%d %s %s\n", lineNumber, now.Format("2006-01-02 15:04:05"), text))
		err := ioutil.WriteFile(fileName, data, 0644)
		if err != nil {
			panic(err)
		}

		lineNumber++
	}
}
