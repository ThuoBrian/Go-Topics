package main

import (
	"fmt"
	"math"
)

func main() {
	s := circle{7.9}
	get_results(s)
}

type circle struct {
	rad float64
}

func (c circle) area() float64 {
	return math.Pi * c.rad * c.rad
}

type shape interface {
	area() float64
}

func get_results(i shape) {
	fmt.Println(i.area())
}
