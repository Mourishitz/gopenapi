package main

import (
	"io"
	"os"
	"path/filepath"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func ReadFromFile(file *os.File) []byte {
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return content
}

func NewFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func ScanDirectory(dirPath string) ([]string, error) {
	var paths []string
	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".yml" {
			paths = append(paths, path)
		}
		return nil
	})
	return paths, err
}
