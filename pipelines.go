package goiterpatterns

import (
	"iter"
)

// https://go.dev/blog/pipelines

func Generate(nums ...int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, n := range nums {
			if !yield(n) {
				break
			}
		}
	}
}

func Square(in iter.Seq[int]) iter.Seq[int] {
	return func(yield func(int) bool) {
		for n := range in {
			if !yield(n * n) {
				break
			}
		}
	}
}
