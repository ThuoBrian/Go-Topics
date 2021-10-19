package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 7, 89}
	sum := 0
	for _, v := range arr {
		sum += v
		fmt.Println("The sum of the slice is: ", sum)
	}
	for i, v := range arr {
		if v == 4 {
			fmt.Println("index:", i)
		}
	}
}
