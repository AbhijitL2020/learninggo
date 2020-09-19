package main

// So many new packages being introduced!
import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var numberFlag = flag.Bool("n", false, "number each line") // This is a valid flag definition

func cat(r *bufio.Reader) {
	i := 1
	for {
		buf, e := r.ReadBytes('\n') // So, open a file from os.Open
		// Then take the file handle and get a bufio Reader for it
		// now read the bytes from the file using this reader line by line
		// And then to check the EOF, use the io package!
		// Hoy, can't that be simpler?!
		// Also that buf string, we need to understand it a bit better
		// I mean is it immutable, whatever that means?
		// Is it better to get allocated each time or declare it once?
		if e == io.EOF {
			break
		}
		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%5d  %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	flag.Parse()          // This fires the parsing, and as we saw weeds out undefined flags!
	if flag.NArg() == 0 { // Number of Arguments. I think these are non-flag arguments only
		cat(bufio.NewReader(os.Stdin)) // Nice touch! It reads Stdin. Echoes of C, Unix heritage!
	}
	for i := 0; i < flag.NArg(); i++ {
		f, e := os.Open(flag.Arg(i))
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n",
				os.Args[0], flag.Arg(i), e.Error()) // Nice to see the error to StdErr
			// Also a good example of OS error handling
			continue
		}
		cat(bufio.NewReader(f))
	}
}
