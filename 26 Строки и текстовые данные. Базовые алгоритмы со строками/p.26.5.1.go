/* Цель задания
Написать программу аналог cat.



Что нужно сделать
Программа должна получать на вход имена двух файлов, необходимо  конкатенировать их содержимое, используя strings.Join.
Если программа запущена командой go run firstFile.txt secondFile.txt resultFile.txt, то она должна написать два соединённых файла в результирующий. */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	// читаем содержимое первого файла
	content1, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// читаем содержимое второго файла
	content2, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// конкатенируем содержимое файлов
	result := strings.Join([]string{string(content1), string(content2)}, "")

	// выводим результат в консоль или записываем в файл
	if len(os.Args) == 3 {
		fmt.Print(result)
	} else {
		err = ioutil.WriteFile(os.Args[3], []byte(result), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
