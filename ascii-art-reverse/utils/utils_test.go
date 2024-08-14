package utils_test

import (
	"ascii-art-reverse/utils"
	"os"
	"testing"
)

func TestParseFile(t *testing.T) {
	// Create a temporary test file
	testFilename := "testfile.txt"
	testData := "Line 1\nLine 2\nLine 3"
	err := os.WriteFile(testFilename, []byte(testData), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer os.Remove(testFilename) // Clean up the test file after test

	// Test the ParseFile function
	lines, err := utils.ParseFile(testFilename)
	if err != nil {
		t.Fatalf("utils.ParseFile returned an error: %v", err)
	}

	// Verify the number of lines
	expectedLines := []string{"Line 1", "Line 2", "Line 3"}
	if len(lines) != len(expectedLines) {
		t.Errorf("expected %d lines, got %d", len(expectedLines), len(lines))
	}

	// Verify the content of each line
	for i, line := range lines {
		if line != expectedLines[i] {
			t.Errorf("expected line %d to be %q, got %q", i+1, expectedLines[i], line)
		}
	}
}

func TestParseFile_FileNotFound(t *testing.T) {
	_, err := utils.ParseFile("nonexistentfile.txt")
	if err == nil {
		t.Error("expected an error when trying to open a nonexistent file, but got nil")
	}
}
