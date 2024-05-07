package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	text_file, err := readFile()
	if err != nil {
		fmt.Println("Error Occured")
		return
	}
	analyzeString(text_file)
}

func readFile() (string, error) {
	filename := "./file/word.txt"
	var file string
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "file not exist", err
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	file = string(data)
	return file, nil
}

func analyzeString(text string) {
	fmt.Println("Analyzing...")

	fmt.Printf("Total Characters: %d\n", len(text))
	fmt.Printf("Total vowels: %d\n", countVowels(text))
	fmt.Printf("Total Words: %d\n", len(strings.Fields(text)))
}

func countVowels(text string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char_text := range text {
		if strings.ContainsRune(vowels, char_text) {
			count++
		}
	}
	return count
}
