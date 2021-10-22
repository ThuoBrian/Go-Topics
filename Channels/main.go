package main

import (
	"fmt"
)

func main() {

	message := make(chan string, 2)

	message <- "This is a buffered channel"
	message <- "This is a second buffered channel"
	msg_1 := <-message
	msg_2 := <-message

	fmt.Println(msg_1)
	fmt.Println(msg_2)

}
