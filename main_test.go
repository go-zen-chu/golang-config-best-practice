package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	originalArgs := os.Args
	defer func () {
		os.Args = originalArgs
	}
	os.Args = []string{}
	main()
}