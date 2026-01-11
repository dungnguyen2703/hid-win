package util

import (
	"os"
	"path/filepath"
)

func GetPath(paths ...string) string {
	return filepath.Join(append([]string{exeDirectory}, paths...)...)
}

func EnsureDir(relativePath string) error {
	path := GetPath(relativePath)
	return os.MkdirAll(path, 0755)
}
