package main

import (
	"flag"
	"log"
	"os"

	"github.com/vncs404/unix-utils/lib/ls"
)

func main() {
	showHiddenFiles := flag.Bool("a", false, "do not ignore entries starting with .")
	directoryIndicator := flag.Bool("p", false, "add a slash to directory names")
	reverseOrder := flag.Bool("r", false, "display entries in reverse order")
	flag.Parse()

	var path string = "."

	args := flag.Args()
	if len(args) > 0 {
		path = args[0]
	}

	config := &ls.LSConfig{
		Path:               path,
		ShowHiddenFiles:    *showHiddenFiles,
		DirectoryIndicator: *directoryIndicator,
		ReverseOrder:       *reverseOrder,
	}

	if err := ls.ListDirFiles(os.DirFS("."), os.Stdout, config); err != nil {
		log.Fatalf("failed to execute 'ls' command, %v", err)
	}
}
