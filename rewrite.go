package main

import (
	"io/ioutil"
	"path/filepath"
)

// rewriteSource rewrites current source and saves
// into temporary file, returning it's path.
func rewriteSource(path string) (string, error) {
	orig, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	data, err := addCode(orig)
	if err != nil {
		return "", err
	}

	tmpDir, err := ioutil.TempDir("", "gotracer_package")
	if err != nil {
		return "", err
	}
	filename := filepath.Join(tmpDir, filepath.Base(path))
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		return "", err
	}

	return tmpDir, nil
}

func addCode(data []byte) ([]byte, error) {
	return data, nil
}
