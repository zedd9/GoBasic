package main

import "fmt"

func main() {
	s := sum(10, 0)
	fmt.Println(s)

	rst := f(10)
	fmt.Println(rst)
}

func sum(x int, s int) int {
	if x == 0 {
		return s
	}

	s += x
	return sum(x-1, s)
}

func f(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}

	return f(x-1) + f(x-2)
}
