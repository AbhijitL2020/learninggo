package main

import "fmt"

func main() {
	var p *int            // p is a pointer to int. In other words, it is going to hold an address
	fmt.Printf("%v\n", p) // It prints nil. In other words, it is initialized to point to nothing

	i := 35
	p = &i // This convention * and & is the same as in C
	fmt.Printf("%v %v \n", p, *p)
	// fmt.Printf("%v \n", *p++)			This gives a compilation error

	// Now for malloc, calloc :), i.e. new. Remember new returns a pointer to an initialized storage
	var q *[]int = new([]int) // q is a pointer to an array of integers. Here initialized to nothing! size?

	// Now, make is slightly different - it returns the variable of type(T), again initialized.
	// An important thing is - it is used for only slices, maps amd channels(whats that?)
	var v []int = make([]int, 5)
	v[0] = 10
	v[1] = 13
	v[2] = 14
	v[3] = 15
	fmt.Printf("As v\n")
	for i = 0; i < 5; i++ {
		fmt.Printf("%d\n", v[i])
	}
	q = &v
	fmt.Printf("As q\n")
	/* for i = 0; i < 5; i++ {   This is not supported. Compiler says indexing q[i] not supported for *[]int
		fmt.Printf("%d", q[i])		That goes to show that such a declaration is pointless. Then how to use it?
	}
	*/
	fmt.Printf("%v\n", q) // The output is: &[10 13 14 15 0]	I get to see this as a pointer to an array/slice
	// This now makes me think of what a slice is.
	// So we know v is a slice by definiton, as we got it using make.
	// can I say?
	/* q = v[2:4]
	fmt.Printf("%v\n", q) */
	// No! It says v is of type []int, and q is of type *[]int

	// So now, I can use it only to pass as a pointer to the array for passing by value
	// Ok lets try that
	fmt.Println("This is what v has to begin with")
	fmt.Printf("%v\n", v)
	// add5(q)		 Again it says: cannot use q (type *[]int) as type []int in argument to add5
	add5(v)
	fmt.Printf("%v\n", q) // output is: &[15 18 19 20 5]  should I infer that array is always pass-by-value?
	// Or is it because I have defined v as a slice? Well lets confirm that.

	var a = [...]int{12, 1, 13, 16, 17}
	fmt.Println("This is what a has to begin with")
	fmt.Printf("%v\n", a)
	// add5(a)  cannot use a (type [5]int) as type []int in function argument: Compiler error!
	// The above makes arrays practically useless! Isn't it?
	// So the below is the workaround - convert array into a slice while passing.
	add5(a[:])
	fmt.Println("This is what a has afterwards")
	fmt.Printf("%v\n", a) // output is: [17 6 18 21 22]  should I infer that SLICE is always pass-by-value?
	/*
		fmt.Println("Send it as a pointer to array")
		add5Ptr(a[:])
		fmt.Println("This is what a has afterwards")
		fmt.Printf("%v\n", a) // output is: [17 6 18 21 22]  should I infer that SLICE is always pass-by-value?
		With all the ingenuity applied to the problem (and my scarce knowledge of Go), it is NO GO!
		One more thought occured to me - why not send q = &a and see what happens? But bored now... :(
	*/
}

func add5(slcQ []int) {
	for i, _ := range slcQ {
		slcQ[i] += 5
		fmt.Printf("%v\n", slcQ[i])
	}
}

/*
func add5Ptr(slcQ *[]int) {
	for i, valueI := range *slcQ {
		valueI += 5
		fmt.Printf("%v\n", valueI)
	}
}
*/
