package storage

import (
	"module_st/pkg/student"
)

func Put(m map[string]*student.Student, s *student.Student) {
	m[s.name] = s
}

func Get(m map[string]*student.Student, name string) *student.Student {
	return m[name]
}
