package ls

import (
	"bytes"
	"fmt"
	"testing"
	"testing/fstest"
)

func TestLS(t *testing.T) {
	t.Run("fs can list files in the passed in directory", func(t *testing.T) {
		fs := fstest.MapFS{
			".bashrc":             {},
			"my_file.txt":         {},
			"my_other_file.md":    {},
			"my-folder/readme.md": {},
			"my-folder/docs.md":   {},
		}

		var tests = []struct {
			Path     string
			Expected string
		}{
			{".", "my-folder  my_file.txt  my_other_file.md\n"},
			{"my-folder", "docs.md  readme.md\n"},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("list the files of %q", test.Path), func(t *testing.T) {
				buffer := bytes.Buffer{}
				err := ListDirFiles(fs, &buffer, &LSConfig{Path: test.Path})

				assertNoError(t, err)
				assertResult(t, buffer.String(), test.Expected)
			})
		}
	})

	t.Run("do not ignore entries starting with a . if configuration is set", func(t *testing.T) {
		fs := fstest.MapFS{
			"my_file.txt":  {},
			".my_dot_file": {},
			".bashrc":      {},
		}

		want := ".bashrc  .my_dot_file  my_file.txt\n"
		buffer := bytes.Buffer{}

		config := &LSConfig{
			Path:            ".",
			ShowHiddenFiles: true,
		}

		err := ListDirFiles(fs, &buffer, config)

		assertNoError(t, err)
		assertResult(t, buffer.String(), want)
	})

	t.Run("append a slash to directories if configuration is set", func(t *testing.T) {
		fs := fstest.MapFS{
			".bashrc":             {},
			"my_file.txt":         {},
			"my-dir/whatever.txt": {},
		}

		want := "my-dir/  my_file.txt\n"
		buffer := bytes.Buffer{}

		config := &LSConfig{
			Path:               ".",
			DirectoryIndicator: true,
		}

		err := ListDirFiles(fs, &buffer, config)

		assertNoError(t, err)
		assertResult(t, buffer.String(), want)
	})

	t.Run("sort file names in reverse alphabetical order", func(t *testing.T) {
		fs := fstest.MapFS{
			"a": {},
			"c": {},
			"d": {},
			"e": {},
			"b": {},
		}

		want := "e  d  c  b  a\n"
		buffer := bytes.Buffer{}

		config := &LSConfig{
			Path:         ".",
			ReverseOrder: true,
		}

		err := ListDirFiles(fs, &buffer, config)

		assertNoError(t, err)
		assertResult(t, buffer.String(), want)
	})
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("got an error but did not expect one, %v", err)
	}
}

func assertResult(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
