// program to find duplicate lines in standard input

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 { // read from std input
        countLines(os.Stdin, counts)
    } else { // read from files
        for _, arg := range(files) {
            f, err := os.Open(arg)
            if err != nil {
                // Fprintf formats according to a format specifier and 
                // writes to w, aka os.Stderr
                fmt.Fprintf(os.Stderr, "Error in opening file, %v\n", err)
                continue
            }
            countLines(f, counts);
            f.Close()
        }
    }
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d, %s\n", n, line)
		}
	}
}

// map is passed by reference
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
