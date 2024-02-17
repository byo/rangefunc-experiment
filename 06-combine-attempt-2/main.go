package main

import (
	"fmt"
	"iter"

	"golang.org/x/exp/constraints"
)

type converter[TIn, TOut any] func(iter.Seq[TIn]) iter.Seq[TOut]

// COMBINE START, OMIT
func combine2[T1, T2 any](
	in iter.Seq[T1],
	l1 converter[T1, T2],
) iter.Seq[T2] {
	return l1(in)
}

func combine3[T1, T2, T3 any](
	in iter.Seq[T1],
	l1 converter[T1, T2],
	l2 converter[T2, T3],
) iter.Seq[T3] {
	return l2(combine2(in, l1))
}

// ... and so on, up to some number n of different converters
// COMBINE END, OMIT

func iota(start, end, step int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i < end; i += step {
			if !yield(i) {
				return
			}
		}
	}
}

func transform[TIn, TOut any](
	t func(TIn) TOut,
) converter[TIn, TOut] {
	return func(input iter.Seq[TIn]) iter.Seq[TOut] {
		return func(yield func(TOut) bool) {
			for val := range input {
				transformed := t(val)
				if !yield(transformed) {
					return
				}
			}
		}
	}
}

func filter[T any](
	f func(T) bool,
) converter[T, T] {
	return func(input iter.Seq[T]) iter.Seq[T] {
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
}

func square[T constraints.Integer](v T) T {
	return v * v
}

func without[T constraints.Integer](withoutValue T) func(T) bool {
	return func(v T) bool { return v != withoutValue }
}

func main() {
	for i := range combine3(
		iota(0, 10, 1),
		transform(square[int]),
		filter(without(16)),
	) {
		fmt.Println(i)
	}
}
