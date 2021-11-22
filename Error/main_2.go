package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-10.45)
	if err != nil {
		log.Fatal(err)
	}

}
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("Invaild, value")

	}
	return 43, nil
}
