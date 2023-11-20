package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	words := flag.Bool("w", false, "Count words")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *words, *bytes))
}

func count(r io.Reader, countLines bool, countWords bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	if countLines {
		scanner.Split(bufio.ScanLines)
	}

	if countWords {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	result := 0
	for scanner.Scan() {
		result++
	}

	return result
}
