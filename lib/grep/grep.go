package grep

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	reset = "\033[0m"
	red   = "\033[31m"
)

func Grep(fs fs.FS, file string, query string, w io.Writer) error {
	f, err := fs.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", file, err)
	}

	scanner := bufio.NewScanner(f)
	buffer := bytes.Buffer{}

	for scanner.Scan() {
		lineContent := scanner.Text()

		if strings.Contains(lineContent, query) {
			fmt.Fprintln(&buffer, strings.Replace(lineContent, query, red+query+reset, -1))
		}
	}

	fmt.Fprint(w, buffer.String())
	return nil
}
