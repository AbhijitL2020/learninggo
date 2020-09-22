//
// Package mygenstack will give multiple implementations of generalized stack
package mygenstack

import (
	"fmt"
	"strconv"
)

const (
	GSTACKFULL  = 10
	GSTACKEMPTY = -1
)

type stacknode struct {
	userVal interface{}
}

type GStack struct {
	stackPtr  int // I would love it to start with -1, rather than 0
	stackData [GSTACKFULL]stacknode
}

// Method New allows me to rest the stackPtr to -1, which indicates GSTACKEMPTY condition
// stype = 0: int, = 1: string
func (s *GStack) New() {
	s.stackPtr = -1
}

// Method Push will add a value at the top of the stack.
// If stackfull, then it returns err
func (s *GStack) Push(v interface{}) bool {
	if s.stackPtr < GSTACKFULL {
		s.stackPtr++
		s.stackData[s.stackPtr].userVal = v
		return true
	}
	return false // in comma ok form, we will get ok to be false, in case of an error
}

// Method Pop will return the value on the top of the stack.
// Except when Stack is Empty, in which case it will return false for the ok parameter
func (s *GStack) Pop() (interface{}, bool) {
	defer s.decrStackPtr() // Comes handy, isn't it?! But it still is kind of indirect. :(
	if s.stackPtr > GSTACKEMPTY {
		return s.stackData[s.stackPtr].userVal, true
	}
	return 0, false // This is that comma ok format of return value, but it does not allow me to send _
}

func (s *GStack) decrStackPtr() {
	s.stackPtr--
}

// Method PrintStack will print the contents of the stack in the order that it will be popped
// so it will be an exported method, and it will return a string, that can be printed with a %s
func (s *GStack) PrintStack() string {
	strOut := ""
	for i := s.stackPtr; i > GSTACKEMPTY; i-- {
		strOut += "[" + strconv.Itoa(i) + ":" + fmt.Sprintf("%v", s.stackData[i].userVal) + "]"
	}
	return strOut
}
