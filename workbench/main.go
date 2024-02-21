package main

import (
	"fmt"
	"iter"

	"golang.org/x/exp/constraints"
)

func transform[T any](in iter.Seq[T], f func(T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range in {
			transformed := f(i)
			if !yield(transformed) {
				return
			}
		}
	}
}

func filter[T any](
	input iter.Seq[T],
	f func(T) bool,
) iter.Seq[T] {
	return func(yield func(T) bool) {
		for val := range input {
			if f(val) {
				if !yield(val) {
					return
				}
			}
		}
	}
}

func without[T constraints.Integer](withoutValue T) func(v T) bool {
	return func(v T) bool { return v != withoutValue }
}

func square[T constraints.Integer](v T) T {
	return v * v
}

func iota(start, end, step int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i < end; i += step {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	for i := range filter(
		transform(
			iota(0, 10, 2),
			square,
		),
		without(16),
	) {
		fmt.Println(i)
	}
}
