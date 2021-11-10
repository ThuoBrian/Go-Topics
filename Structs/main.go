package main

import (
	"fmt"
)

func main() {
	p1 := person{First: "Brian", Last: "Thuo", age: 28}
	p2 := person{First: "John", Last: "Mboya", age: 27}

	fmt.Println(p1.Full_name())
	fmt.Println(p2.Full_name())

}

type person struct {
	First, Last string
	age         int
}

func (p person) Full_name() string {
	return p.First + p.Last
}
