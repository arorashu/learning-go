// a program that print index and value of each of its arguments per line

package main

import (
	"fmt"
	"os"
)

func main() {

	for index, arg := range os.Args[1:] {
		fmt.Print(index)
		fmt.Println(", " + arg)
	}

}
