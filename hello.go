package main

import "fmt" // implements formatted I/O

/* Print a welcome message */
func main() {
	fmt.Println("Hello, World!")
	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range list {
		// do something with k and v
		fmt.Println(k, v)
	}
	for pos, char := range "Gő!" {
		fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
	}
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The third element is %d\n", arr[2]) // It did not give a null ptr err, but 0
	indx := arr[1] + (arr[0] / 10)
	fmt.Printf("The seventeenth element is %d\n", arr[indx]) // It gave compilation error for 17
	// And when it was a calculation and therefore runtime error, it gave a panic
}
