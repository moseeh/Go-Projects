package utils

import (
	"testing"
)

func TestColor(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		color     string
		expected  string
		wantError bool
	}{
		{"Red color", "Hello", "red", "\033[38;2;255;0;0mHello\033[0m", false},
		{"Green color", "World", "green", "\033[38;2;0;255;0mWorld\033[0m", false},
		{"Blue color", "Test", "blue", "\033[38;2;0;0;255mTest\033[0m", false},
		{"Color not found", "Oops", "unknown", "COLOR NOT FOUND: ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Color(tt.input, tt.color)
			if tt.wantError {
				if result != tt.expected {
					t.Errorf("expected an error message, got %v", result)
				}
			} else {
				if result != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}
