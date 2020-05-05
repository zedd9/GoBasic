package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) PrintName() {
	fmt.Print(p.name)
}

func main() {
	p1 := Person{"개똥이", 15}
	p2 := Person{name: "소똥이", age: 15}
	p3 := Person{name: "Carson"}
	p4 := Person{}

	var p Person
	p.name = "Smith"
	p.age = 24

	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)

	p.PrintName()
}
