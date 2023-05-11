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
