package main

import (
	"fmt"
)

func main() {

	bol, even := half(32)
	fmt.Println(bol, even)

}
func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

//Question:
//Write a function which takes an integer and returns two values:
//the integer divided by 2
//whether or not the integer is even (true, false)
