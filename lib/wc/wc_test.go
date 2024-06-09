package wc

import (
	"bytes"
	"testing"
	"testing/fstest"
)

func TestWC(t *testing.T) {
	t.Run("returns correct file data", func(t *testing.T) {
		fileContent := "This is the content of my file\nIt has two lines"
		fs := fstest.MapFS{
			"file.txt": {Data: []byte(fileContent)},
		}

		want := "2 11 46 file.txt\n"

		buffer := bytes.Buffer{}
		err := WordCount(fs, "file.txt", &buffer)

		assertNoError(t, err)
		assertResult(t, buffer.String(), want)
	})

	t.Run("returns an error if folder is passed in", func(t *testing.T) {
		fs := fstest.MapFS{
			"file.txt":            {},
			"my-dir/whatever.txt": {},
		}

		buffer := bytes.Buffer{}
		err := WordCount(fs, "my-dir", &buffer)

		assertError(t, err)
	})
}

func assertError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Error("expected an error but didn't get one")
	}
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
