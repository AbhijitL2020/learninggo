/* Exercise 1 from Learning Go */
package main

import (
	"fmt"
	"strconv"
) // implements formatted I/O

func main() {
	for i := 1; i <= 25; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			output("FizzBuzz")
		} else if i%3 == 0 {
			output("Fizz")
		} else if i%5 == 0 {
			output("buzz")
		} else {
			output(strconv.Itoa(i))
		}
	}
}

func output(strOut string) {
	// fmt.Printf("%s\n", strOut)
	fmt.Println(strOut)
}
