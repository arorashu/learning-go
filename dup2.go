// program to find duplicate lines in standard input

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]map[string]int)
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
            input := bufio.NewScanner(f)
            for input.Scan() {
                if counts[input.Text()] == nil {
                   counts[input.Text()] = make(map[string]int)
                }
                counts[input.Text()][arg]++
            }
            f.Close()
        }
    }

    for line, occurFiles := range counts {
        count := 0
        var keys []string
        for filename, n := range occurFiles {
            count+=n
            keys = append(keys, filename)
        }
		if count > 1 {
			fmt.Printf("%d, %s, files: %v \n", count, line, keys)
		}
	}
}

// map is passed by reference
func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
        if counts[input.Text()] == nil {
           counts[input.Text()] = make(map[string]int)
        }
        counts[input.Text()]["stdin"]++
	}
}
