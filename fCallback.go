/*
This brings home that I do not understand call back.
My code might be invoking a function. But is it calling back?
What does callback means?

I think it means - hey give me a function to call.

*/
package main

import "fmt"

func main() {
	offset := 3
	for i := 0; i < 10; i++ {
		fCaller(i, offset, plusTwo)
	}
}

func fCaller(i int, offset int, f func(int, int)) {
	f(i, offset)
}

func plusTwo(i int, offset int) {
	fmt.Printf("%d\n", i+offset)
}
