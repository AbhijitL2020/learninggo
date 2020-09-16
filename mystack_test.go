// Package mystack testing: This does not work. I may check this in, but only as a reference
// I think I need a better understanding of what is meant by unit testing, as well as the specifics of
// dealing with a struct and a ptr to struct in the code.
package mystack

import (
	"mystack"
	"testing"
)

func TestPush(t *testing.T) {
	testStk := new(mystack.Stack)
	testStk.New()
	testStk.Push(101)
	testStk.Push(102)
	testStk.Push(103)
	testStk.Push(104)
	testStk.Push(105)
	testStk.Push(106)
	testStk.Push(107)
	testStk.Push(108)
	testStk.Push(109)
	testStk.Push(110)
	testStk.Push(111)

	for {
		v, ok = testStk.Pop()
		if ok != true {
			t.Log("It failed")
			t.Fail()
		}
		// t.Fatal("not implemented")
	}
}
