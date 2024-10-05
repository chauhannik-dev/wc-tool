package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func countCharacters(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening the file")
		os.Exit(1)
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}

	totalChars := len(content)

	return totalChars, nil
}

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}

func countWords(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	wordsCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		wordsCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordsCount, nil
}

func main() {
	// Define flags
	countFlag := flag.Bool("c", false, "count characters flag")
	linesFlag := flag.Bool("l", false, "count lines flag")
	wordsFlag := flag.Bool("w", false, "count words flag")

	flag.Parse() // parses the value of the flag

	args := flag.Args() // parses the non flag values, like the name of the txt file

	if len(args) < 1 {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}

	filename := args[0]

	if *countFlag {
		totalChars, err := countCharacters(filename)
		if err != nil {
			fmt.Println("Error counting the characters: ", err)
			os.Exit(1)
		}

		fmt.Println(totalChars, filename)
	}

	if *linesFlag {
		lines, err := countLines(filename)
		if err != nil {
			fmt.Println("Error counting the lines: ", err)
			os.Exit(1)
		}

		fmt.Println(lines, filename)
	}

	if *wordsFlag {
		words, err := countWords(filename)
		if err != nil {
			fmt.Println("Error counting the words: ", err)
			os.Exit(1)
		}

		fmt.Println(words, filename)
	}

}
