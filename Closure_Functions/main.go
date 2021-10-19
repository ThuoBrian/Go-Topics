package main

import "fmt"

func main() {
	newInt := new_itm()

	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())

	New_Ints := new_itm()
	fmt.Println(New_Ints())
}

func new_itm() func() int {
	i := 0
	return func() int {
		i++
		return (i)
	}

}
