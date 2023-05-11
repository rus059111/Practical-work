package student

import (
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent(s string) *Student {
	fields := strings.Fields(s)
	age, _ := strconv.Atoi(fields[1])
	grade, _ := strconv.Atoi(fields[2])
	return &Student{Name: fields[0], Age: age, Grade: grade}
}
