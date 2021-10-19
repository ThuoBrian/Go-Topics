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
	i := 0
	for i = 0; i < 10; i++ {
		fmt.Println("The values in the fibonacci", fib(i))
	}
}
