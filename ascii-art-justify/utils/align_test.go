package utils

import (
	"fmt"
	"testing"
)

func TestAlign(t *testing.T) {
	type args struct {
		align        string
		asciiArt     string
		input        string
		contentLines []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Left Align",
			args: args{
				align:    "left",
				asciiArt: "Hello\nWorld",
				input:    "Hello World",
				contentLines: []string{
					"Hello",
					"World",
				},
			},
		},
		{
			name: "Right Align",
			args: args{
				align:    "right",
				asciiArt: "Hello\nWorld",
				input:    "Hello World",
				contentLines: []string{
					"Hello",
					"World",
				},
			},
		},
		{
			name: "Center Align",
			args: args{
				align:    "center",
				asciiArt: "Hello\nWorld",
				input:    "Hello World",
				contentLines: []string{
					"Hello",
					"World",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fmt.Printf("Running test: %s\n", test.name)
			Align(test.args.align, test.args.asciiArt, test.args.input, test.args.contentLines)
		})
	}
}
