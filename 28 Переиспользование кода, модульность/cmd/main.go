/* package main

import (
	"fmt"
)

type Student struct {
	name  string
	age   int
	grade int
}

func newStudent(name string, age int, grade int) *Student {
	return &Student{
		name:  name,
		age:   age,
		grade: grade,
	}
}

func main() {
	students := make(map[string]*Student)

	for {
		var name string
		var age, grade int
		_, err := fmt.Scanln(&name, &age, &grade)
		if err != nil {
			break
		}

		student := newStudent(name, age, grade)
		students[name] = student
	}

	fmt.Println("Студенты из хранилища:")
	for name, student := range students {
		fmt.Println(name, student.age, student.grade)
	}
}
*/

/* Цель задания

Научиться декомпозировать и рефакторить код на примере программы, написанной в прошлом уроке.


Что нужно сделать

В прошлом домашнем задании вы написали программу для работы со студентами. Мы указываем в стандартный ввод данные о студенте, а после получения сигнала об окончании работы программы она выводит имена всех студентов на экран.

Вам нужно отрефакторить код прошлого домашнего задания. Декомпозируйте его так, чтобы логике одной сущности соответствовал один пакет.

Для того, чтобы вы могли использовать методы и переменные, которые объявлены в других пакетах, сделайте их экспортируемыми.

Структура программы после рефакторинга может выглядеть следующим образом:


Критерии оценки

Зачёт:

    программа работает в прежнем режиме, но декомпозирована на пакеты;
    при получении одной строки (например, «имяСтудента 24 1») программа создаёт данные о студенте и сохраняет их, далее ожидает следующую строку или сигнал EOF (Сtrl + d);
    при получении сигнала EOF программа должна вывести имена всех студентов из map.


На доработку:

    задание не выполнено или выполнено не полностью.


Как отправить задание на проверку

Выполните задание в файле вашей среды разработки и пришлите ссылку на архив с вашим проектом через форму ниже.
*/

package main

import (
	"bufio"
	"fmt"
	"module_st/pkg/storage"
	"module_st/pkg/student"
	"os"
)

func main() {
	m := make(map[string]*student.Student)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := student.NewStudent(scanner.Text())
		storage.Put(m, s)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Считывание ввода:", err)
	}
	fmt.Println("Студенты из хранилища:")
	for name := range m {
		fmt.Printf("%s %d %d\n", m[name].name, m[name].age, m[name].grade)
	}
}
