package models

import (
	"os"
	"time"
)

// infomation about the flag
type FileInfo struct {
	Name      string
	Mode      os.FileMode
	Size      int64
	ModTime   time.Time
	IsDir     bool
	User      string
	Group     string
	Links     int
	BlockSize int64
}

// flags to use
type Options struct {
	Long       bool
	Recursive  bool
	All        bool
	Reverse    bool
	SortByTime bool
}
