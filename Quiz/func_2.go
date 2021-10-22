package main

import "fmt"

func main() {
	largest := max(4, 5, 3, 2, 2, 1, 4, 66, 6, 2, 21)
	fmt.Println("the biggest value is: ", largest)
}
func max(n ...int) int {
	greatest := 0
	for _, v := range n {
		if v > greatest {
			greatest = v
		}
	}
	return greatest
}

//Question
//Write a function with one variadic parameter that finds the greatest number in a list of numbers
