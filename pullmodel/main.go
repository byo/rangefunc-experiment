package main

import (
	"fmt"
	"iter"
)

func generator(yield func(int) bool) {
	for i := 0; yield(i); i++ {
	}
}

func main() {
	next, close := iter.Pull(generator)
	defer close()

	for i, ok := next(); ok && i < 1000; i, ok = next() {
		fmt.Println(i)
	}
}
