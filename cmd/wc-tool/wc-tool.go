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

func getStats(file *os.File) (int64, int, int, error) {
	bytes, err := countBytes(file)
	if err != nil {
		return 0, 0, 0, err
	}

	lines, err := countLines(file)
	if err != nil {
		return 0, 0, 0, err
	}

	// Reset file pointer after it scans the file once for counting lines
	// So reset the pointer from the end of the file to the top of it
	_, err = file.Seek(0, 0)
	if err != nil {
		return 0, 0, 0, err
	}

	words, err := countWords(file)
	if err != nil {
		return 0, 0, 0, err
	}

	return bytes, lines, words, nil
}

func main() {
	// Define flags
	countFlag := flag.Bool("c", false, "count bytes flag")
	linesFlag := flag.Bool("l", false, "count lines flag")
	wordsFlag := flag.Bool("w", false, "count words flag")

	flag.Parse() // parses the value of the flag

	args := flag.Args() // parses the non flag values, like the name of the txt file

	var arg *os.File
	var filename string

	if len(args) > 0 {
		filename = args[0]

		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}

		arg = file

		defer file.Close()
	} else {
		file, err := os.CreateTemp("", "temp-*.txt")

		if err != nil {
			fmt.Println("Error creating temp file:", err)
			os.Exit(1)
		}
		defer os.Remove(file.Name())

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			_, err := file.WriteString(scanner.Text() + "\n")
			if err != nil {
				fmt.Println("Error writing to temp file:", err)
				os.Exit(1)
			}
		}

		file.Sync()

		file, err = os.Open(file.Name())
		if err != nil {
			fmt.Println("Error reopening temp file:", err)
			os.Exit(1)
		}

		arg = file
		filename = file.Name()

		defer file.Close()
	}

	if flag.NFlag() == 0 {
		chars, lines, words, err := getStats(arg)
		if err != nil {
			fmt.Println("Problem in calculating the stats: ", err)
			os.Exit(1)
		}

		fmt.Printf("Lines: %d | Words: %d | Characters: %d ", lines, words, chars)
	}

	if *countFlag {
		totalChars, err := countBytes(arg)
		if err != nil {
			fmt.Println("Error counting the characters: ", err)
			os.Exit(1)
		}

		fmt.Println("Total characters in the file: ", totalChars)
	}

	if *linesFlag {
		lines, err := countLines(arg)
		if err != nil {
			fmt.Println("Error counting the lines: ", err)
			os.Exit(1)
		}

		fmt.Println("Total lines in the file: ", lines)
	}

	if *wordsFlag {
		words, err := countWords(arg)
		if err != nil {
			fmt.Println("Error counting the words: ", err)
			os.Exit(1)
		}

		fmt.Println("Total words in the file: ", words)
	}

}
