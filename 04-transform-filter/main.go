package main

import (
	"fmt"
	"iter"

	"golang.org/x/exp/constraints"
)

func iota(start, end, step int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i < end; i += step {
			if !yield(i) {
				return
			}
		}
	}
}

func transform[T any](
	input iter.Seq[T], // Take some other sequence as an input
	t func(T) T,
) iter.Seq[T] {
	return func(yield func(T) bool) {
		for val := range input {
			transformed := t(val)
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

// HELPERS START, OMIT

func square[T constraints.Integer](v T) T {
	return v * v
}

func without[T constraints.Integer](withoutValue T) func(v T) bool {
	return func(v T) bool { return v != withoutValue }
}

// HELPERS END, OMIT

func main() {
	for i := range filter(
		transform(
			iota(0, 10, 1),
			square,
		),
		without(16),
	) {
		fmt.Println(i)
	}
}
