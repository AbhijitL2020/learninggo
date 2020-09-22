// Program to try out Generalized Stack again
package main

import (
	"fmt"
	"mygenstack"
	"reflect"
)

type person struct {
	id    int
	fname string
	lname string
}

func main() {
	errCount := 0
	fmt.Println("GenStack")
	stk := new(mygenstack.GStack) // So this something that I did not get to read in the book.

	p1 := person{1, "Abhijit", "Laghate"}
	p2 := new(person)

	fmt.Printf("type of p1 = %v & type of p2 = %v\n", reflect.TypeOf(p1), reflect.TypeOf(p2))
	// tname := reflect.TypeOf(p1)
	// tname1 := "person"
	// p3 := new(tname1)
	// p3 := new(reflect.TypeOf(p1))
	// fmt.Printf("type of p3 = %v \n", reflect.TypeOf(p3))

	// This is for the Int version fo the stack
	stk.New() // I am happy that my New() function did work!
	stk.Push(8)
	if !stk.Push("Hello") { // Not true in 2: This should fail and provide an error and it does
		errCount++
	}
	stk.Push(p1)
	stk.Push(10)
	stk.Push(7)
	fmt.Println(stk.PrintStack())
	/*
		// This is the string version of the stack
		stk.New(1) // I am happy that my New() function did work!
		stk.Push("hi")
		if !stk.Push(9) { // This should fail and provide an error and it does
			errCount++
		}
		stk.Push("hello")
		stk.Push("bye")
		fmt.Println(stk.PrintStack())
	*/
	for {
		v, ok := stk.Pop()
		if ok == false { // This was master-stroke!
			break
		}
		fmt.Printf("%v %v\n", v, ok)
	}
	fmt.Println(errCount)
}
