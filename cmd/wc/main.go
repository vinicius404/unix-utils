package main

import (
	"log"
	"os"

	"github.com/vncs404/unix-utils/lib/wc"
)

func main() {
	var fileName string

	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalf("file name was not provided")
	} else {
		fileName = args[0]
	}

	if err := wc.WordCount(os.DirFS("."), fileName, os.Stdout); err != nil {
		log.Fatalf("failed to execute 'ws' command: %v", err)
	}
}
