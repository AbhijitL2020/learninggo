/* Exercise 1 from Learning Go */
package main

import "fmt" // implements formatted I/O

func main() {
	for i := 1; i <= 25; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Printf("%d FizzBuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d Buzz\n", i)
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
