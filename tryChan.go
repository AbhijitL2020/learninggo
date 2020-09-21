package main

import "fmt"

func main() {
	ch := make(chan int)
	// quit := make(chan bool)
	quitint := make(chan int) // Added to check the blocking/non-blocking nature. It worked!
	// To test the earlier code, change every quitint to quit, and make it a chan bool
	go shower(ch, quitint) // how reasonable is it to use two channels for such a pidling task?
	for i := 0; i < 10; i++ {
		ch <- i // So this is a channel between the instance of shower running and main
	}
	quitint <- 0             // or true, does not matter	-- Until the main signals end, via sending a message thr channel quit
	fmt.Println("Quit main") // IMHO, if quit main prints after quit shower, it shows that chan waited to be read
	// IMHO, when I change the chanel to int, it is unbuffered, hence non-blocking
	// So when 0 is written to quitint, the main continues, without blocking, and the program finished double quick
	// It does not gie a chance for shower to finish!
}

func shower(c chan int, quitint chan int) {
	for {
		select {
		case j := <-c:
			fmt.Printf("%d\n", j) // If you have read a value on channel c, then print it
		case <-quitint:
			fmt.Println("QuitInt shower")
			break // Once you rcvd a value on channel quit, you quit
		}
	}
}
