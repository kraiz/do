package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	doFileName = "do.yml"
)

func main() {
	fmt.Println(os.Args[1:])

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't detect current working directory: %s", err)
	}

	dofile, err := findDoFile(cwd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "do.yml not found: %s", err)
	}

	fmt.Printf("Using %s.\n\n", dofile)
}

func findDoFile(path string) (string, error) {
	candidate := filepath.Join(path, doFileName)
	if _, err := os.Stat(candidate); os.IsNotExist(err) {
		parent := filepath.Dir(path)
		if parent == path {
			return "", fmt.Errorf("%s not found", doFileName)
		}
		return findDoFile(parent)
	}
	return candidate, nil
}
