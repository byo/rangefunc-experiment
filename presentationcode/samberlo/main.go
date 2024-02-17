package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	result := lo.Filter(
		lo.Map(
			numbers,
			func(i int, _ int) int { return i * i },
		),
		func(i int, _ int) bool { return i != 16 },
	)

	for _, i := range result {
		fmt.Println(i)
	}
}
