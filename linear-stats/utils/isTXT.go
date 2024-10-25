package utils

import "path/filepath"

func IsTxt(filename string) bool {
	extentsion := filepath.Ext(filename)
	return extentsion == ".txt"
}
