package main

import (
	"log"
	"os"

	"github.com/vncs404/unix-utils/lib/cat"
)

func main() {
	var fileName string

	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalf("file name was not provided")
	} else {
		fileName = args[0]
	}

	if err := cat.Catenate(os.DirFS("."), fileName, os.Stdout); err != nil {
		log.Fatalf("failed to execute 'cat' command, %v", err)
	}
}
