package utils

import (
	"os"
	"testing"
)

func TestModeToString(t *testing.T) {
	testCases := []struct {
		name     string
		mode     os.FileMode
		expected string
	}{
		{
			name:     "Directory with full permissions",
			mode:     os.ModeDir | 0o777,
			expected: "drwxrwxrwx",
		},
		{
			name:     "Regular file with read/write permissions",
			mode:     0o644,
			expected: "-rw-r--r--",
		},
		{
			name:     "Symbolic link",
			mode:     os.ModeSymlink | 0o777,
			expected: "lrwxrwxrwx",
		},
		{
			name:     "Block device with read-only permissions",
			mode:     os.ModeDevice | 0o400,
			expected: "br--------",
		},
		{
			name:     "Character device with specific permissions",
			mode:     os.ModeDevice | os.ModeCharDevice | 0o744, // Changed permissions to match `r--r--r-x`
			expected: "crwxr--r--",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := modeToString(tc.mode)
			if result != tc.expected {
				t.Errorf("modeToString(%v) = %v, want %v", tc.mode, result, tc.expected)
			}
		})
	}
}
