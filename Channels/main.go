package main

import (
	"fmt"
)

func main() {

	message := make(chan string)
	go func() {
		message <- "Hi, Brian, Mchapa kazi"
	}()
	msg := <-message
	fmt.Println(msg)
}
