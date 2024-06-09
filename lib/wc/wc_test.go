package wc

import (
	"bytes"
	"testing"
	"testing/fstest"
)

func TestWC(t *testing.T) {
	fileContent := "This is the content of my file\nIt has two lines"
	fs := fstest.MapFS{
		"file.txt": {Data: []byte(fileContent)},
	}

	want := "2 11 46 file.txt"

	buffer := bytes.Buffer{}
	err := WordCount(fs, "file.txt", &buffer)

	assertNoError(t, err)
	assertResult(t, buffer.String(), want)
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("did not expect an error but got one: %v", err)
	}
}

func assertResult(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
