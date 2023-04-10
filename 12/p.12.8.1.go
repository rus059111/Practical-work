package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Открываем файл для записи
	file, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineNumber := 1

	for {

		fmt.Print("Введите строку: ")
		var Text string
		fmt.Scanln(&Text)
		now := time.Now()
		if Text == "exit" {
			break
		}
		s := fmt.Sprintf("%d %s %s\n", lineNumber, now.Format("2006-01-02 15:04:05"), Text)

		file.WriteString(s)

		lineNumber++
	}
}
