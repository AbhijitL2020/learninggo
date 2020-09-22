// Package mygenstack is my implementation of a GENERIC stack using struct
// This uses interfacees. And whatever I have learnt at this point.
// Objective is: User uses only one type to deal with stack. She does not have to worry about whether int or string
// Once the intent is decalred at the start, then usage remains identical
//
// NOTE: THIS IS A FAILED ATTEMPT. I DO NOT LIKE TO WASTE MEMORY LIKE THAT. 20200920
//
package mygenstack

import "strconv"

const (
	STACKFULL  = 10
	STACKEMPTY = -1
)

// try01: Did not work. It need type of nil !:)
// var stackData = nil // So it is a pointer set to nil

type Stack struct {
	stackPtr  int // I would love it to start with -1, rather than 0
	stackType int
	// stackData [STACKFULL]int // This will anyway be filled up
	stackIntData [STACKFULL]int
	stackStrData [STACKFULL]string
}

// Method New allows me to rest the stackPtr to -1, which indicates STACKEMPTY condition
// stype = 0: int, = 1: string
func (s *Stack) New(stype int) {
	s.stackPtr = -1
	/*
		if stype == 0 {
			s.stackIntData = make([STACKFULL]int)
		}
		if stype == 1 {
			s.stackStrData = make([STACKFULL]string)
		}
	*/
	s.stackType = stype
}

// Method Push will add a value at the top of the stack.
// If stackfull, then it returns err
func (s *Stack) Push(v interface{}) bool {
	switch v.(type) {
	case string:
		if s.stackType == 1 {
			goto LabelDoomPush
		} else {
			return false
		}
	case int:
		if s.stackType == 0 {
			goto LabelDoomPush
		} else {
			return false
		}
	}
LabelDoomPush: // This is horrible code! You should have just let the code fall through!!
	if s.stackPtr < STACKFULL {
		s.stackPtr++
		switch s.stackType {
		case 0:
			s.stackIntData[s.stackPtr] = v.(int) // Does this work?! :)
		case 1:
			s.stackStrData[s.stackPtr] = v.(string)
		}
		return true
	}
	return false // in comma ok form, we will get ok to be false, in case of an error
}

// Method Pop will return the value on the top of the stack.
// Except when Stack is Empty, in which case it will return false for the ok parameter
func (s *Stack) Pop() (interface{}, bool) {
	defer s.decrStackPtr() // Comes handy, isn't it?! But it still is kind of indirect. :(
	if s.stackPtr > STACKEMPTY {
		switch s.stackType {
		case 0:
			return s.stackIntData[s.stackPtr], true
		case 1:
			return s.stackStrData[s.stackPtr], true
		}
	}
	return 0, false // This is that comma ok format of return value, but it does not allow me to send _
}

func (s *Stack) decrStackPtr() {
	s.stackPtr--
}

// Method PrintStack will print the contents of the stack in the order that it will be popped
// so it will be an exported method, and it will return a string, that can be printed with a %s
func (s *Stack) PrintStack() string {
	strOut := ""
	switch {
	case s.stackType == 0:
		for i := s.stackPtr; i > STACKEMPTY; i-- {
			strOut += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.stackIntData[i]) + "]"
		}
	case s.stackType == 1:
		for i := s.stackPtr; i > STACKEMPTY; i-- {
			strOut += "[" + strconv.Itoa(i) + ":" + s.stackStrData[i] + "]"
		}
	}
	return strOut
}
