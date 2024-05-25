package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Write a function that takes a filename and return the content of the file
func ReadFile(filename string, n int, isByte bool) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	if n > 0 {
		if isByte {
			b := make([]byte, n)
			_, err := file.Read(b)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return ""
			}
			return string(b)
		} else {
			scanner := bufio.NewScanner(file)
			var lines []string
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
				if len(lines) == n {
					break
				}
			}
			return strings.Join(lines, "\n")
		}
	}
	// Get the file size
	stat, _ := file.Stat()
	size := stat.Size()

	// Read the content of the file
	content := make([]byte, size)
	file.Read(content)

	return string(content)
}
