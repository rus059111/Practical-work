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

package main

import (
	"bufio"
	"fmt"
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
