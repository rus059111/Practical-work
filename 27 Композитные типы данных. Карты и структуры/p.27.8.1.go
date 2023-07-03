/* Цель задания

Научиться работать с композитными типами данных: структурами и картами


Что нужно сделать

Напишите программу, которая считывает ввод с stdin, создаёт структуру student и записывает указатель на структуру в хранилище map[studentName] *Student.

type Student struct {

name string

age int

grade int

}

Программа должна получать строки в бесконечном цикле, создать структуру Student через функцию newStudent, далее сохранить указатель на эту структуру в map, а после получения EOF (ctrl + d) вывести на экран имена всех студентов из хранилища. Также необходимо реализовать методы put, get.


Общие условия

Разработка выполняется в среде golang или vs code.

Input


go run main.go

Строки


Вася 24 1

Семен 32 2

EOF


Output


Студенты из хранилища:

Вася 24 1

Семен 32 2


Критерии оценки

Зачёт:

    при получении одной строки (например, «имяСтудента 24 1») программа создаёт студента и сохраняет его, далее ожидает следующую строку или сигнал EOF (Сtrl + Z);
    при получении сигнала EOF программа должна вывести имена всех студентов из map.

На доработку:

    задание не выполнено или выполнено не полностью.


Как отправить задание на проверку

Выполните задание в файле Go, загрузите файл через онлайн-редактор REPL или в виде архива и пришлите ссылку на документ через форму ниже. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

func newStudent(s string) *Student {
	fields := strings.Fields(s)
	age, _ := strconv.Atoi(fields[1])
	grade, _ := strconv.Atoi(fields[2])
	return &Student{name: fields[0], age: age, grade: grade}
}

func put(m map[string]*Student, s *Student) {
	m[s.name] = s
}

func get(m map[string]*Student, name string) *Student {
	return m[name]
}

func main() {
	m := make(map[string]*Student)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := newStudent(scanner.Text())
		put(m, s)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Считывание ввода:", err)
	}
	fmt.Println("Студенты из хранилища:")
	for name := range m {
		fmt.Printf("%s %d %d\n", m[name].name, m[name].age, m[name].grade)
	}
}
