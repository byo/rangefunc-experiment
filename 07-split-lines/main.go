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
	input iter.Seq[T],
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

func square[T constraints.Integer](v T) T {
	return v * v
}

func without[T constraints.Integer](withoutValue T) func(v T) bool {
	return func(v T) bool { return v != withoutValue }
}

func main() {
	seq1 := iota(0, 10, 1)
	seq2 := transform(seq1, square)
	seq3 := filter(seq2, without(16))
	for i := range seq3 {
		fmt.Println(i)
	}
}
