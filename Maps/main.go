package main

import "fmt"

func main() {
	var Mygreeting = make(map[string]string)

	Mygreeting["Brian"] = "Good Morning"
	Mygreeting["Thuo"] = "Habari za asubuhi"

	fmt.Println(Mygreeting)
}
