package main

import "fmt"

func iota(
	yield func(int) bool, // Ignoring return value for now
) {
	yield(1)
	yield(2)
	yield(3)
}

func main() {
	for i := range iota {
		fmt.Println(i)
	}
}
