package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	// read input from args
	if len(files) == 0 {
		// if empty args, use stdin
		countLines(os.Stdin, counts)
	} else {
		// else open and read each file passed as arg
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			// prints 'nDuplicates		line'
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(readFrom *os.File, counts map[string]int) {
	input := bufio.NewScanner(readFrom)
	for input.Scan() {
		// adds lines as part of a map with a count
		counts[input.Text()]++
	}
	if input.Err() != nil {
		panic(input.Err())
	}
}
