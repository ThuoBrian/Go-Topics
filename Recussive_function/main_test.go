package main

import (
	"testing"
)

func main() {
}
func TestFib(t *testing.T) {
	fibonacii := fib()
	if fibonacii != 60 {
		t.Error("The Fibonacii has an error")

	}
}
