package main

import (
	"log"
	"os"
)

func main() {
	// TODO input validation
	err := os.Symlink(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}
