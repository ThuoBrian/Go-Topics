package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log_err.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("on-file.txt")
	if err != nil {
		log.Println("Error Happened", err)

	}

}
