/* Цель задания

Написать программу аналог cat.


Что нужно сделать

Программа должна получать на вход имена двух файлов, необходимо  конкатенировать их содержимое, используя strings.Join.


Что оценивается

    При получении одного файла на входе программа должна печатать его содержимое на экран.
    При получении двух файлов на входе программа соединяет их и печатает содержимое обоих файлов на экран.
    Если программа запущена командой go run firstFile.txt secondFile.txt resultFile.txt, то она должна написать два соединённых файла в результирующий.


Общие условия

Разработка выполняется в среде golang или vs code.

first.txt


контент первого файла

second.txt


контент второго файла

result .txt


контент первого файла

контент второго файла

Input


go run first.txt second.txt result.txt


first.txt


контент первого файла

second.txt


контент второго файла

Input


go run first.txt second.txt

Output


контент первого файла

контент второго файла


Как отправить задание на проверку

Выполните задание в файле Go. Загрузите файл на Google Диск, откройте доступ для всех по ссылке. Отправьте ссылку на файл через форму для сдачи домашнего задания.

Или отправьте файл через онлайн-редактор REPL, или архивом. */
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
