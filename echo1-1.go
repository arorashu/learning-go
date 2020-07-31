// a simple echo program that prints its arguments, and the name of the command that invoked it

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
