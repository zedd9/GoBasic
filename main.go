package main

import "fmt"

func main() {
	a := 9
	b := 4

	fmt.Printf("a=%v b=%v\n", a, b)
	fmt.Printf("a&b = %v\n", a&b)
	fmt.Printf("a|b = %v\n", a|b)
	fmt.Printf("a^b = %v\n", a^b)
	fmt.Printf("a%%b = %v\n", a%b)
	fmt.Printf("a&^a = %v\n", a&^a)
	fmt.Printf("a<<1 = %v\n", a<<1)
	fmt.Printf("a>>1 = %v\n", a>>1)
}
