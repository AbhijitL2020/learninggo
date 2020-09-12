package main

import "fmt"

func main() {
	fmt.Println(chkPanic(ferr))
}

// Now I need to write a function that panics
func ferr() {
	var a [4]int
	a[3] = 15
}

// Now we need to write a function that takes input of another function and reports it it panicked
func chkPanic(f func()) (b bool) {
	defer func() {
		if r := recover(); r != nil {
			b = true
		}
	}() // this call to the function is most important, as defer needs to define and call a function
	f()
	return false
}
