package main

import (
	"fmt"
)

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x)

}
func zero(z *int) {
	fmt.Println(z)
	*z = 0
}
