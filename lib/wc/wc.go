package wc

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

func WordCount(fs fs.FS, file string, w io.Writer) error {
	f, err := fs.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", file, err)
	}

	stat, err := f.Stat()
	if err != nil {
		return fmt.Errorf("failed to retrieve file information %s: %v", file, err)
	}

	if stat.IsDir() {
		return fmt.Errorf("%s is a directory", file)
	}

	lineCount := 0
	wordCount := 0
	byteCount := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		content := scanner.Text()
		byteCount += len([]byte(content))
		wordCount += len(strings.Fields(content))
		lineCount++
	}

	fmt.Fprintf(w, "%d %d %d %s\n", lineCount, wordCount, byteCount, file)
	return nil
}
