package main

import (
	"bufio"
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
	if err := os.Chmod("log.txt", 04444); err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(file)
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

		writer.WriteString(s)

		lineNumber++
	}
}
