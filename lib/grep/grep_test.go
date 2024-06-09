package grep

import (
	"bytes"
	"testing"
	"testing/fstest"
)

func TestGrep(t *testing.T) {
	t.Run("returns line where query was found", func(t *testing.T) {
		content := "Welcome to my world.\nYou will be greeted by the unexpected here and your mind will be challenged\nand expanded in ways you never tought possible."
		fs := fstest.MapFS{
			"file.txt": {Data: []byte(content)},
		}

		query := "greeted"
		want := "You will be greeted by the unexpected here and your mind will be challenged\n"
		buffer := bytes.Buffer{}

		Grep(fs, "file.txt", query, &buffer)

		assertResult(t, buffer.String(), want)
	})

	t.Run("returns multiple lines where the query was found", func(t *testing.T) {
		content := "One dollar and eighty-seven cents.\nThat was all.\nAnd sixty cents of it was in pennies."
		fs := fstest.MapFS{
			"file.txt": {Data: []byte(content)},
		}

		query := "cents"
		want := "One dollar and eighty-seven cents.\nAnd sixty cents of it was in pennies.\n"
		buffer := bytes.Buffer{}

		Grep(fs, "file.txt", query, &buffer)

		assertResult(t, buffer.String(), want)
	})

	t.Run("returns nothing if no query was found", func(t *testing.T) {
		content := "One dollar and eighty-seven cents.\nThat was all.\nAnd sixty cents of it was in pennies."
		fs := fstest.MapFS{
			"file.txt": {Data: []byte(content)},
		}

		query := "vegetables"
		want := ""
		buffer := bytes.Buffer{}

		Grep(fs, "file.txt", query, &buffer)

		assertResult(t, buffer.String(), want)
	})
}

func assertResult(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
