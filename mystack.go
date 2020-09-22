// Package mystack is my implementation of a stack using struct
// This is the package for it.
package mystack

import "strconv"

const (
	STACKFULL  = 10
	STACKEMPTY = -1
)

type Stack struct {
	stackPtr  int            // I would love it to start with -1, rather than 0
	stackData [STACKFULL]int // This will anyway be filled up
}

// Method New allows me to rest the stackPtr to -1, which indicates STACKEMPTY condition
func (s *Stack) New() {
	s.stackPtr = -1
}

// Method Push will add a value at the top of the stack.
// If stackfull, then it returns err
func (s *Stack) Push(v int) bool {
	if s.stackPtr < STACKFULL {
		s.stackPtr++
		s.stackData[s.stackPtr] = v
		return true
	}
	return false // in comma ok form, we will get ok to be false, in case of an error
}

// Method Pop will return the value on the top of the stack.
// Except when Stack is Empty, in which case it will return false for the ok parameter
func (s *Stack) Pop() (int, bool) {
	if s.stackPtr > STACKEMPTY {
		v := s.stackData[s.stackPtr]
		s.stackPtr--
		return v, true
	}
	return 0, false // This is that comma ok format of return value, but it does not allow me to send _
}

// Method PrintStack will print the contents of the stack in the order that it will be popped
// so it will be an exported method, and it will return a string, that can be printed with a %s
func (s *Stack) PrintStack() string {
	strOut := ""
	for i := s.stackPtr; i > STACKEMPTY; i-- {
		strOut += "[" + strconv.Itoa(i) + ":" + strconv.Atoi(s.stackData[i]) + "]"
	}
	return strOut
}
