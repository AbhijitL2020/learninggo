// Program tryMyDblList is used to create a doubly linked list
package main

import (
	"fmt"
	"mydbllist"
)

func main() {
	l := new(mydbllist.MyDblList)
	l.New()
	l.Add(10)
	l.Add(12)
	l.Add(14)
	l.Add(8)
	fmt.Println(l.PrintAll())
}
