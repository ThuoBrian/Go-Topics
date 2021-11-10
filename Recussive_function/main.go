package main

import (
	"fmt"
)

func fib(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func main() {
	sum := 0
	var x [12]int
	i := 0
	for i = 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(fib(i))
		}
		x[i] = i
	}
	for _, v := range x {
		sum += v
	}
	fmt.Println("the sum of the fibonacci numbers", sum)
}
