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
	// indx := arr[1] + (arr[0] / 10)  // Remember, if it is declared and not used, it is an error!
	//	fmt.Printf("The seventeenth element is %d\n", arr[indx]) // It gave compilation error for 17
	// And when it was a calculation and therefore runtime error, it gave a panic

	//20200910 - Slices

	var array [100]int
	slice := array[0:99]
	slice[98] = 1
	// slice[99] = 2	// This generates a run time error
	fmt.Printf("slice[98]  %d\n", slice[98])

	var name string
	name = "Abhijit Laghate"
	c := []rune(name)
	for k, v := range c {
		fmt.Printf("%d  - %c\n", k, v)
	}
	sfname := c[0:7]
	slname := c[8:] // This allows us to just leave it to the end of the array
	for _, v := range sfname {
		fmt.Printf("%c", v)
	}
	fmt.Printf("\n")
	for _, v := range slname { // did you see the cleverness here? :)
		fmt.Printf("%c", v)
	}
	fmt.Printf("\n")
	// No but this will be even more clever...
	fmt.Printf("%v", slname) // It prints... [76 97 103 104 97 116 101] :( But god try I guess
	fmt.Printf("\n")
}
