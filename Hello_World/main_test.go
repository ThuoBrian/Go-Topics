package main

import (
	"testing"
)

func TestHi(t *testing.T) {
	greeting := hi()
	if greeting != "hello world" {
		t.Error("The code should be Hello World")
	}
}
