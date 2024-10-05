package main

import (
	"flag"
	"fmt"
	"io"
	"os"
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

func main() {
	// Define flags
	countFlag := flag.Bool("c", false, "count characters in the flag")

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

}
