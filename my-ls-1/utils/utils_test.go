package utils_test

import (
	"reflect"
	"testing"

	"my-ls-1/models"
	"my-ls-1/utils" // Replace "your_module_name" with the actual module name
)

func TestSortPath(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Already sorted",
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Reverse sorted",
			input:    []string{"c", "b", "a"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Mixed order",
			input:    []string{"b", "a", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "Single element",
			input:    []string{"a"},
			expected: []string{"a"},
		},
		{
			name:     "Duplicate elements",
			input:    []string{"b", "a", "b", "c"},
			expected: []string{"a", "b", "b", "c"},
		},
		{
			name:     "Case sensitive",
			input:    []string{"B", "a", "C", "b"},
			expected: []string{"B", "C", "a", "b"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the input slice to avoid modifying the original
			input := make([]string, len(tc.input))
			copy(input, tc.input)

			utils.SortPath(input)

			if !reflect.DeepEqual(input, tc.expected) {
				t.Errorf("SortPath() = %v, want %v", input, tc.expected)
			}
		})
	}
}

func TestParseArgs(t *testing.T) {
	testCases := []struct {
		name            string
		args            []string
		expectedOptions models.Options
		expectedPaths   []string
	}{
		{
			name:            "No args",
			args:            []string{},
			expectedOptions: models.Options{},
			expectedPaths:   []string{},
		},
		{
			name:            "Single short flag",
			args:            []string{"-l"},
			expectedOptions: models.Options{Long: true},
			expectedPaths:   []string{},
		},
		{
			name:            "Multiple short flags",
			args:            []string{"-la"},
			expectedOptions: models.Options{Long: true, All: true},
			expectedPaths:   []string{},
		},
		{
			name:            "Long flag",
			args:            []string{"--recursive"},
			expectedOptions: models.Options{Recursive: true},
			expectedPaths:   []string{},
		},
		{
			name:            "Mixed flags and paths",
			args:            []string{"-l", "path1", "--recursive", "path2", "-a"},
			expectedOptions: models.Options{Long: true, Recursive: true, All: true},
			expectedPaths:   []string{"path1", "path2"},
		},
		{
			name:            "All options",
			args:            []string{"-l", "-R", "-a", "-r", "-t", "path1"},
			expectedOptions: models.Options{Long: true, Recursive: true, All: true, Reverse: true, SortByTime: true},
			expectedPaths:   []string{"path1"},
		},
		{
			name:            "Invalid flag treated as path",
			args:            []string{"-x", "path1"},
			expectedOptions: models.Options{},
			expectedPaths:   []string{"-x", "path1"},
		},
		{
			name:            "Combined short flags",
			args:            []string{"-laRrt", "path1"},
			expectedOptions: models.Options{Long: true, All: true, Recursive: true, Reverse: true, SortByTime: true},
			expectedPaths:   []string{"path1"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			options, paths := utils.ParseArgs(tc.args)

			if !reflect.DeepEqual(options, tc.expectedOptions) {
				t.Errorf("ParseArgs() options = %v, want %v", options, tc.expectedOptions)
			}

			if !reflect.DeepEqual(paths, tc.expectedPaths) {
				t.Errorf("ParseArgs() paths = %v, want %v", paths, tc.expectedPaths)
			}
		})
	}
}
