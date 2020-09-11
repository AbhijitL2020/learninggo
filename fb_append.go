/* Using a different algo */
package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i, s := 1, ""; i <= 25; i, s = i+1, "" {
		if i%3 == 0 {
			s = "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		} else {
			if s == "" {
				s = strconv.Itoa(i)
			}
		}
		fmt.Println(s)
	}
}
