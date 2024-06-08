package cat

import (
	"bytes"
	"testing"
	"testing/fstest"
)

func TestCat(t *testing.T) {
	content := "These are the contents of my file\nLine breaks and all\n"
	fs := fstest.MapFS{
		"file.txt": {Data: []byte(content)},
	}

	buffer := bytes.Buffer{}

	err := Catenate(fs, "file.txt", &buffer)

	assertNoError(t, err)
	assertResult(t, buffer.String(), content)
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
