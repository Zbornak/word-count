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
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *words))
}

func count(r io.Reader, countLines bool, countWords bool) int {
	scanner := bufio.NewScanner(r)
	if countLines {
		scanner.Split(bufio.ScanLines)
	}

	if countWords {
		scanner.Split(bufio.ScanWords)
	}

	result := 0
	for scanner.Scan() {
		result++
	}

	return result
}
