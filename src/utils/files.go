package utils

import (
	"bufio"
	"os"
)

// FileExists checks whether a file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// WriteStringInFile writes a string to a given file
func WriteStringInFile(path string, content string) {
	file, err := os.Create(path)
	Check(err)
	defer file.Close()

	w := bufio.NewWriter(file)
	_, err = w.WriteString(content)
	Check(err)
	w.Flush()
}
