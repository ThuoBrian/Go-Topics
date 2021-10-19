package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is an even number")
	} else {
		fmt.Println("7 is an odd Number")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisibe by 4")
	}
	if num := 9; num < 0 {
		fmt.Println(num, "Is a negative value")
	} else if num <= 10 {
		fmt.Println(num, "Has one digit")
	} else {
		fmt.Println(num, "has multiple values")
	}
}
