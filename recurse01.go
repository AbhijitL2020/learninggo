package main

import "fmt"

func main() {
	fmt.Println("First recursive function, in a long time...")
	frec(10)
}

func frec(lim int) {
	if lim == 0 {
		return
	}
	fmt.Printf("%d\n", lim)
	frec(lim - 1)
	fmt.Printf("%d\n", lim)
}
