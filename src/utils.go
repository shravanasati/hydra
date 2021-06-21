package main

import (
	"fmt"
	"os"
)

// makeFile creates a file with the provided content.
func makeFile(filename, content string) {
	f, e := os.Create(filename)
	handleException(e)
	_, er := f.WriteString(content)
	handleException(er)
	defer f.Close()
	cwd, _ := os.Getwd()
	fmt.Printf("\n - Created file '%v' at %v.", filename, cwd)
}
