package utils

import "testing"

func TestIsEnglish(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Hello, World!", true},    // English sentence
		{"123", true},              // Numbers are also considered English
		{"¡Hola, mundo!", false},   // Spanish characters
		{"こんにちは、世界！", false},       // Japanese characters
		{"Привет, мир!", false},    // Russian characters
		{"", true},                 // Empty string
		{"\n\t\r", false},          // Special characters
		{"\x7F", false},            // Non-printable ASCII character
		{"😊", false},               // Emoji
		{"Hello, World! 😀", false}, // English with emoji
		{"Hello, 世界!", false},      // English with non-English characters
		{"123 456", true},          // English with numbers and spaces
	}

	for _, test := range tests {
		if output := IsEnglish(test.input); output != test.expected {
			t.Errorf("IsEnglish(%q) = %v, expected %v", test.input, output, test.expected)
		}
	}
}
