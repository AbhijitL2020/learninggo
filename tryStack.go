package main

import (
	"fmt"
	"mystack"
)

func main() {
	fmt.Println("Stack")
	stk := new(mystack.Stack) // So this something that I did not get to read in the book.
	stk.New()                 // I am happy that my New() function did work!
	stk.Push(8)
	stk.Push(10)
	stk.Push(7)
	for {
		v, ok := stk.Pop()
		if ok == false { // This was master-stroke!
			break
		}
		fmt.Printf("%d %v\n", v, ok)
	}
}
