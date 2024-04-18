package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const WIDTH = 10

func main() {
	// Parsing the flags
	// c := flag.Bool("c", false, "print the byte counts")
	// flag.Parse()

	// Get only the files from the args
	var files []string
	for _, v := range os.Args {
		if checkIfFileExists(v) {
			files = append(files, v)
		}
	}

	var total int64
	for _, v := range files {
		bytes := fileBytes(v)
		total += bytes
		fmt.Printf("%*d %s\n", WIDTH, bytes, v)
	}
	fmt.Printf("%*d total\n", WIDTH, total)
}

func checkIfFileExists(path string) bool {
	if strings.Contains(path, "wcg") {
		return false
	}
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func fileBytes(path string) int64 {
	f, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return f.Size()
}
