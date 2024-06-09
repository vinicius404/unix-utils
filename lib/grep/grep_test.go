package grep

import (
	"bytes"
	"fmt"
	"testing"
	"testing/fstest"
)

func TestGrep(t *testing.T) {
	t.Run("return correct lines for the specified query", func(t *testing.T) {
		cases := []struct {
			content  string
			query    string
			expected string
			filename string
		}{
			{
				content:  "Welcome to my world.\nYou will be greeted by the unexpected here and your mind will be challenged\nand expanded in ways you never tought possible.",
				query:    "greeted",
				expected: fmt.Sprintf("You will be %sgreeted%s by the unexpected here and your mind will be challenged\n", red, reset),
				filename: "file_1.txt",
			},
			{
				content:  "One dollar and eighty-seven cents.\nThat was all.\nAnd sixty cents of it was in pennies.",
				query:    "cents",
				expected: fmt.Sprintf("One dollar and eighty-seven %scents%s.\nAnd sixty %scents%s of it was in pennies.\n", red, reset, red, reset),
				filename: "file_2.txt",
			},
			{
				content:  "One dollar and eighty-seven cents.\nThat was all.\nAnd sixty cents of it was in pennies.",
				query:    "vegetables",
				expected: "",
				filename: "file_3.txt",
			},
		}

		fs := fstest.MapFS{}

		for _, c := range cases {
			fs[c.filename] = &fstest.MapFile{Data: []byte(c.content)}
		}

		for _, c := range cases {
			t.Run(fmt.Sprintf("return correct lines for %q with %q query", c.filename, c.query), func(t *testing.T) {
				buffer := bytes.Buffer{}
				err := Grep(fs, c.filename, c.query, &buffer)

				assertNoError(t, err)
				assertResult(t, buffer.String(), c.expected)
			})
		}
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
