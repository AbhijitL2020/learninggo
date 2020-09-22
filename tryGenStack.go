package main

import (
	"fmt"
	"mygenstack"
)

func main() {
	errCount := 0
	fmt.Println("GenStack")
	stk := new(mygenstack.Stack) // So this something that I did not get to read in the book.
	/*
		// This is for the Int version fo the stack
			stk.New(0) // I am happy that my New() function did work!
			stk.Push(8)
			stk.Push(10)
			stk.Push(7)
			fmt.Println(stk.PrintStack())
	*/
	// This is the string version of the stack
	stk.New(1) // I am happy that my New() function did work!
	stk.Push("hi")
	if !stk.Push(9) { // This should fail and provide an error and it does
		errCount++
	}
	stk.Push("hello")
	stk.Push("bye")
	fmt.Println(stk.PrintStack())

	for {
		v, ok := stk.Pop()
		if ok == false { // This was master-stroke!
			break
		}
		fmt.Printf("%v %v\n", v, ok)
	}
	fmt.Println(errCount)
}
