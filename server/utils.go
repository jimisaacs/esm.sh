package server

import (
	"os"
	"strings"
)

const (
	// EOL defines the char of end of line
	EOL = "\n"
)

func isValidatedESImportPath(importPath string) bool {
	return strings.HasPrefix(importPath, "/") || strings.HasPrefix(importPath, "./") || strings.HasPrefix(importPath, "../")
}

func startsWith(s string, prefixs ...string) bool {
	for _, prefix := range prefixs {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

func endsWith(s string, suffixs ...string) bool {
	for _, suffix := range suffixs {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

func ensureExt(path string, ext string) string {
	if !strings.HasSuffix(path, ext) {
		return path + ext
	}
	return path
}

func ensureDir(dir string) (err error) {
	_, err = os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
	}
	return
}

func fileExists(filepath string) bool {
	fi, err := os.Lstat(filepath)
	return err == nil && !fi.IsDir()
}