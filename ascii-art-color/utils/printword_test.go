package utils

import (
	"io"
	"os"
	"testing"
)

func TestPrintWord(t *testing.T) {
	contentLine, err := os.ReadFile("standard.txt")
	if err != nil {
		t.Fatalf("failed to read standard.txt: %v", err)
	}
	contentLines := SplitFile(string(contentLine))

	tests := []struct {
		name     string
		word     string
		expected string
	}{
		{"Single word", "Hello", " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			PrintWord(test.word, contentLines)

			w.Close()
			output, _ := io.ReadAll(r)
			os.Stdout = oldStdout

			if string(output) != test.expected {
				t.Errorf("PrintWord(%q) = %q, expected %q", test.word, output, test.expected)
			}
		})
	}
}
