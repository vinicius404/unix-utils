package grep

import (
	"bytes"
	"fmt"
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
		want := fmt.Sprintf("You will be %sgreeted%s by the unexpected here and your mind will be challenged\n", Red, Reset)
		buffer := bytes.Buffer{}

		err := Grep(fs, "file.txt", query, &buffer)

		assertNoError(t, err)
		assertResult(t, buffer.String(), want)
	})

	t.Run("returns multiple lines where the query was found", func(t *testing.T) {
		content := "One dollar and eighty-seven cents.\nThat was all.\nAnd sixty cents of it was in pennies."
		fs := fstest.MapFS{
			"file.txt": {Data: []byte(content)},
		}

		query := "cents"
		want := fmt.Sprintf("One dollar and eighty-seven %scents%s.\nAnd sixty %scents%s of it was in pennies.\n", Red, Reset, Red, Reset)
		buffer := bytes.Buffer{}

		err := Grep(fs, "file.txt", query, &buffer)

		assertNoError(t, err)
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

		err := Grep(fs, "file.txt", query, &buffer)

		assertNoError(t, err)
		assertResult(t, buffer.String(), want)
	})
}

func assertResult(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("got an error but did not expect one, %v", err)
	}
}
