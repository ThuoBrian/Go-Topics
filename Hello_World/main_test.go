package main

import (
	"testing"
)

func TestHi(t *testing.T) {
	greeting := hi()
	if greeting != "hello world" {
		t.Error("Please check this code function")
	}
}
