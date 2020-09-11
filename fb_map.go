/* FizzBuzz with maps */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	const (
		FIZZ     = "3"
		IFIZZ    = 3
		BUZZ     = "5"
		IBUZZ    = 5
		FIZZBUZZ = "FizzBuzz"
	)
	fizzbuzz := map[string]string{
		FIZZ: "Fizz", BUZZ: "Buzz",
	}
	for i := 1; i <= 25; i++ {
		if (i%IFIZZ == 0) && (i%IBUZZ == 0) {
			output(FIZZBUZZ)
		} else if i%IFIZZ == 0 {
			output(fizzbuzz[FIZZ])
		} else if i%IBUZZ == 0 {
			output(fizzbuzz[BUZZ])
		} else {
			output(strconv.Itoa(i))
		}
	}
}

func output(strOut string) {
	// fmt.Printf("%s\n", strOut)
	fmt.Println(strOut)
}
