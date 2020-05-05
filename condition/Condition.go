package main

import "fmt"

func main() {
	var a bool
	a = 3 < 4 && 2 < 5

	fmt.Println("3 < 4 && 2 < 5 = ", a)

	fmt.Print("3 > 4 ? ")
	if 3 > 4 {
		fmt.Println("참")
	} else {
		fmt.Println("거짓")
	}

	b := 5
	if b == 3 {
		fmt.Println("b는 3")
	} else if b == 4 {
		fmt.Println("b는 4")
	} else {
		fmt.Println("b는 3도 아니고 4도 아니다.")
	}
	fmt.Println("b의 값은 5")
}
