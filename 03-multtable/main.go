package main

import "fmt"

func iota(start, end, step int) func(yield func(int) bool) {
	return func(yield func(int) bool) {
		for i := start; i < end; i += step {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	for i := range iota(1, 11, 1) {
		for j := range iota(1, 11, 1) {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}
}
