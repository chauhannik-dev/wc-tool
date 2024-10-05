package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func countBytes(file *os.File) (int64, error) {
	info, err := file.Stat()

	if err != nil {
		return 0, err
	}

	size := info.Size()

	return size, nil
}

func countLines(file *os.File) (int, error) {
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

func countWords(file *os.File) (int, error) {
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
	countFlag := flag.Bool("c", false, "count bytes flag")
	linesFlag := flag.Bool("l", false, "count lines flag")
	wordsFlag := flag.Bool("w", false, "count words flag")

	flag.Parse() // parses the value of the flag

	args := flag.Args() // parses the non flag values, like the name of the txt file

	if len(args) < 1 {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}

	filename := args[0]

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening a file")
		os.Exit(1)
	}

	defer file.Close()

	if *countFlag {
		totalChars, err := countBytes(file)
		if err != nil {
			fmt.Println("Error counting the characters: ", err)
			os.Exit(1)
		}

		fmt.Println(totalChars, filename)
	}

	if *linesFlag {
		lines, err := countLines(file)
		if err != nil {
			fmt.Println("Error counting the lines: ", err)
			os.Exit(1)
		}

		fmt.Println(lines, filename)
	}

	if *wordsFlag {
		words, err := countWords(file)
		if err != nil {
			fmt.Println("Error counting the words: ", err)
			os.Exit(1)
		}

		fmt.Println(words, filename)
	}

}
