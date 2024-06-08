package cat

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
)

// Scan a file line by line and send to stdout.
func Catenate(fs fs.FS, file string, w io.Writer) error {
	buffer := bytes.Buffer{}

	f, err := fs.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", file, err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Fprintln(&buffer, scanner.Text())
	}

	fmt.Fprint(w, buffer.String())
	return nil
}
