// a simple echo program that prints its arguments

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// 1. standard assignments
	start := time.Now()
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += arg + sep
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("method 1 time:", time.Since(start).Seconds())

	// 2. using strings join
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("method 2 time:", time.Since(start).Seconds())

}
