package main

import "fmt"

func main() {

	//setup the pipeline
	c := gen(5, 6)
	out := sq(c)

	fmt.Println(<-out)
	fmt.Println(<-out)
}
func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
func sq(ins chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range ins {
			out <- n * n
		}
		close(out)
	}()
	return out
}
