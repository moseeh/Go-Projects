package asciiArt

import (
	"os"
	"testing"
)

func TestBannerFile(t *testing.T) {
	// Test with "standard" banner
	os.Args = []string{"cmd", "input", "standard"}
	expected := "banners/standard.txt"
	if result := BannerFile(); result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}

	// Test with "shadow" banner
	os.Args = []string{"cmd", "input", "shadow"}
	expected = "banners/shadow.txt"
	if result := BannerFile(); result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}

	// Test with "thinkertoy" banner
	os.Args = []string{"cmd", "input", "thinkertoy"}
	expected = "banners/thinkertoy.txt"
	if result := BannerFile(); result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}

	// Test with an unsupported banner (should return an empty string)
	os.Args = []string{"cmd", "input", "unsupported"}
	expected = ""
	if result := BannerFile(); result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}

	// Test with only one argument (should default to "standard")
	os.Args = []string{"cmd", "input"}
	expected = "banners/standard.txt"
	if result := BannerFile(); result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}

	// Test with no arguments (should return an empty string)
	os.Args = []string{"cmd"}
	expected = ""
	if result := BannerFile(); result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
