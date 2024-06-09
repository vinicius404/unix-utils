package main

import (
	"log"
	"os"

	"github.com/vncs404/unix-utils/lib/grep"
)

func main() {
	var query string
	var fileName string

	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatalf("Usage: grep [QUERY] [FILE]")
	} else {
		query = args[0]
		fileName = args[1]
	}

	if err := grep.Grep(os.DirFS("."), fileName, query, os.Stdout); err != nil {
		log.Fatalf("failed to execute 'cat' command: %v", err)
	}
}
