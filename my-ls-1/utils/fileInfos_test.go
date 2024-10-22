package utils

import (
	"io/fs"
	"os"
	"testing"
	"time"

	"my-ls-1/models"
)

// mockDirEntry implements fs.DirEntry and wraps mockFileInfo.
type mockDirEntry struct {
	fileInfo mockFileInfo
}

func (m mockDirEntry) Name() string               { return m.fileInfo.name }
func (m mockDirEntry) IsDir() bool                { return m.fileInfo.isDir }
func (m mockDirEntry) Type() fs.FileMode          { return m.fileInfo.Mode() }
func (m mockDirEntry) Info() (fs.FileInfo, error) { return m.fileInfo, nil }

// mockFileInfo implements os.FileInfo.
type mockFileInfo struct {
	name  string
	isDir bool
}

func (m mockFileInfo) Name() string       { return m.name }
func (m mockFileInfo) Size() int64        { return 1024 }
func (m mockFileInfo) Mode() fs.FileMode  { return 0o644 }
func (m mockFileInfo) ModTime() time.Time { return time.Now() }
func (m mockFileInfo) IsDir() bool        { return m.isDir }
func (m mockFileInfo) Sys() interface{}   { return nil }

func TestGetFileInfos(t *testing.T) {
	// Mock directory entries using mockFileInfo
	entries := []fs.DirEntry{
		mockDirEntry{fileInfo: mockFileInfo{name: "file1.txt", isDir: false}},
		mockDirEntry{fileInfo: mockFileInfo{name: "dir1", isDir: true}},
		mockDirEntry{fileInfo: mockFileInfo{name: ".hidden", isDir: false}},
	}

	testCases := []struct {
		name     string
		path     string
		entries  []fs.DirEntry
		options  models.Options
		expected int
	}{
		{
			name:     "No options",
			path:     "testdir",
			entries:  entries,
			options:  models.Options{},
			expected: 2,
		},
		{
			name:     "All files",
			path:     "testdir",
			entries:  entries,
			options:  models.Options{All: true},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getFileInfos(tc.path, tc.entries, tc.options)
			if len(result) != tc.expected {
				t.Errorf("getFileInfos() returned %d items, expected %d", len(result), tc.expected)
			}
		})
	}
}

func TestGetFileInfo(t *testing.T) {
	now := time.Now()
	info := mockFileInfo{name: "testfile", isDir: false}

	testCases := []struct {
		name     string
		fileInfo os.FileInfo
		options  models.Options
		expected models.FileInfo
	}{
		{
			name:     "Basic file info",
			fileInfo: info,
			options:  models.Options{},
			expected: models.FileInfo{
				Name:    "testfile",
				Mode:    0o644,
				Size:    1024,
				ModTime: now,
				IsDir:   false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getFileInfo(tc.fileInfo.Name(), tc.fileInfo, tc.options)
			if result.Name != tc.expected.Name ||
				result.Mode != tc.expected.Mode ||
				result.Size != tc.expected.Size ||
				result.IsDir != tc.expected.IsDir {
				t.Errorf("GetFileInfo() = %v, want %v", result, tc.expected)
			}
		})
	}
}
