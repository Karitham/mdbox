package main

import "path/filepath"

func elems(s string, els []string) bool {
	for _, e := range els {
		if s == e {
			return true
		}
	}
	return false
}

func cleanRelPath(pwd string, path string) string {
	path = filepath.ToSlash(path)
	if nsrc, err := filepath.Rel(pwd, path); err == nil {
		return nsrc
	}
	return path
}
