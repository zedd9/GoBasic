package main

import "fmt"

type Student struct {
	name  string
	class int

	sungjuk Sungjuk
}

type Sungjuk struct {
	name  string
	grade string
}

func (s Student) viewSungjuk() {
	fmt.Print(s.sungjuk)
}

func (s *Student) inputSungjuk(name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

func main() {
	var s Student
	s.name = "철수"
	s.class = 1

	s.sungjuk.name = "수학"
	s.sungjuk.grade = "A"

	s.viewSungjuk()

	s.inputSungjuk("과학", "B")
	s.viewSungjuk()
}
