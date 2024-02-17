package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
	}()

	ch2 := make(chan int)
	go func() {
		defer close(ch2)
		for i := range ch1 {
			ch2 <- i * i
		}
	}()

	// ... (continued on the next slide)
	// SPLIT-POINT OMIT
	// ... (continued)

	ch3 := make(chan int)
	go func() {
		defer close(ch3)
		for i := range ch2 {
			if i != 16 {
				ch3 <- i
			}
		}
	}()

	for i := range ch3 {
		fmt.Println(i)
	}
}
