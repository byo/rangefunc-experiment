package main

import (
	"fmt"
	"io"
	"iter"
	"os"
)

func open(fname string) iter.Seq[io.Reader] {
	return func(yield func(io.Reader) bool) {
		fl, err := os.Open(fname)
		if err != nil {
			fmt.Printf("Error opening file: %v", err)
			return
		}

		yield(fl) // Yield a single file only

		fl.Close() // And cleanup after it was used
	}
}

func main() {
	for fl := range open("main.go") {
		io.Copy(os.Stdout, fl)
	}
}
