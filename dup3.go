// program to find duplicate lines in filename in args
// read file completely, NOT in streaming mode

package main

import (
	"fmt"
	"os"
    "io/ioutil"
    "strings"
)

func main() {

	counts := make(map[string]int)
    for _, filename := range(os.Args[1:]) {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
                fmt.Fprintf(os.Stderr, "Error in reading file, %v\n", err)
                continue
        }
        for _, line := range(strings.Split(string(data), "\n") ) {
            counts[line]++
        }

    }
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d, %s\n", n, line)
		}
	}
}

