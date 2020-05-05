package main

import "fmt"

func main() {
	s := ("Hello 월드")

	fmt.Printf("len(s) = %d\n", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], ", ")
	}

	s2 := []rune(s)
	fmt.Printf("len(s2) = %d\n", len(s2))
	for i := 0; i < len(s2); i++ {
		fmt.Print(s2[i], ", ")
	}
}
