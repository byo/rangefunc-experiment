package main

import (
	"fmt"
	"iter"

	"golang.org/x/exp/constraints"
)

func square[T constraints.Integer](v T) T {
	return v * v
}

func without[T constraints.Integer](withoutValue T) func(T) bool {
	return func(v T) bool { return v != withoutValue }
}

type combinable[T any] iter.Seq[T]

func iota(start, end, step int) combinable[int] {
	return func(yield func(int) bool) {
		for i := start; i < end; i += step {
			if !yield(i) {
				return
			}
		}
	}
}

func (input combinable[T]) transform(
	t func(T) T,
) combinable[T] {
	return func(yield func(T) bool) {
		for val := range input {
			transformed := t(val)
			if !yield(transformed) {
				return
			}
		}
	}
}

func (input combinable[T]) filter(
	f func(T) bool,
) func(yield func(T) bool) {
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

func main() {
	for i := range iota(0, 10, 1).
		transform(square).
		filter(without(16)) {
		fmt.Println(i)
	}
}
