package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 20; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}
func bar() {
	for j := 0; j < 20; j++ {
		fmt.Println("Bar:", j)
	}
	wg.Done()
}
