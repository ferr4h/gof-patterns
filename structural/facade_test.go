package structural

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureStdout(f func(string), name string) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f(name)
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = old
	return buf.String()
}

func TestFacade(t *testing.T) {
	facade := NewHomeTheater()

	output := captureStdout(facade.WatchMovie, "Into the wild")
	lines := strings.Split(output, "\n")

	expectedActions := []string{
		"Screen is down",
		"Disk has been inserted",
		"Lights are off",
		"Movie Into the wild is playing",
	}

	for i, action := range expectedActions {
		if action != lines[i] {
			t.Errorf("\nExpected action: %s\nGot: %s", action, lines[i])
		}
	}
}
