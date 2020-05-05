package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var i int
	for {
		if i == 5 {
			i++
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
		i++
	}
}
