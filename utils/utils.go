package utils

import (
	"log"
	"os"
	"strings"
)

func ReadFileToLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\n")
	var cleanLines []string
	for _, line := range lines {
		var cleanedLine string = strings.TrimSpace(line)
		cleanLines = append(cleanLines, cleanedLine)
	}
	return cleanLines
}
