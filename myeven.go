// This is the first one where we are including a package
// It is a simple program which determines whether the given number is even or odd
package main

import (
	"even"
	"fmt"
)

func main() {

	i := 5
	fmt.Printf("%d even? %v\n", i, even.Even(i))
}
