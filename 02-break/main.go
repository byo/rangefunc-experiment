package main

import "fmt"

func iota(yield func(int) bool) {
	for i := 0; ; i++ {
		if !yield(i) {
			return
		}
	}
}

func main() {
	for i := range iota {
		if i == 1000 {
			break // yield -> return false
		}
		if i%2 == 0 {
			continue // yield -> return true
		}
		fmt.Println(i)
		// yield -> return true
	}
}
