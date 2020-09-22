// Package myDblList will implement a doubly linked list
//
package mydbllist

import "strconv"

type node struct { // This is a internal strcuture. The world does not need to know the impl
	nodeVal int
	next    *node
	prev    *node
}

type MyDblList struct { // This is the one the world will use
	ptrFirstNode   *node // Note while the struct is public, the members are private
	ptrLastNode    *node
	ptrCurrentNode *node // Initially current node will be nil, and point to nothing!
	CntNodes       int   // This is only one that is public .Inititally this will be 0
}

func (l *MyDblList) New() { // so we initialized it.
	l.CntNodes = 0
	l.ptrFirstNode = nil
	l.ptrLastNode = nil
	l.ptrCurrentNode = nil
}

func (l *MyDblList) Add(v int) bool { // It will be a method Add of a variable m of type pointer to myDblList
	// It will take input of an integer
	// It will return an indication of success or failure
	// allocate the
	m := new(node)
	m.nodeVal = v
	if l.ptrLastNode == nil { // we can safely assume that this is the first node
		l.ptrFirstNode, l.ptrCurrentNode, l.ptrLastNode = m, m, m
	} else {
		(l.ptrLastNode).next = m // Set the next node of the last node
		m.prev = l.ptrLastNode   // set the prev of the new node to the last node
		m.next = nil             // set the next of the new node to nil
		l.ptrLastNode = m        // And this new node is now the last node
	}
	l.CntNodes++
	return true
}

// PrintAll is used for formatted printing of the doubly-linked list
func (l *MyDblList) PrintAll() (strOut string) {
	if l.ptrFirstNode == nil {
		strOut = "[:] No Contents"
		return
	}
	i := 0
	for c := l.ptrFirstNode; c != nil; c = c.next {
		strOut += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(c.nodeVal) + "]"
	}
	strOut += "\nIndex: " + strconv.Itoa(i)
	strOut += "Count: " + strconv.Itoa(l.CntNodes)
	return
}
