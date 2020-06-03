package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestNextPermutation(_t *testing.T) {
	i, o := []int{1, 2, 3}, []int{1, 3, 2}
	input_str := S(i)
	nextPermutation(i)
	T(_t, input_str, S(i), S(o))

	i, o = []int{1, 3, 2}, []int{2, 1, 3}
	input_str = S(i)
	nextPermutation(i)
	T(_t, input_str, S(i), S(o))

	i, o = []int{3, 2, 1}, []int{1, 2, 3}
	input_str = S(i)
	nextPermutation(i)
	T(_t, input_str, S(i), S(o))

	i, o = []int{1, 1, 5}, []int{1, 5, 1}
	input_str = S(i)
	nextPermutation(i)
	T(_t, input_str, S(i), S(o))

	i, o = []int{2, 3, 1}, []int{3, 1, 2}
	input_str = S(i)
	nextPermutation(i)
	T(_t, input_str, S(i), S(o))

	i, o = []int{1, 2, 3, 2, 1}, []int{1, 3, 1, 2, 2}
	input_str = S(i)
	nextPermutation(i)
	T(_t, input_str, S(i), S(o))

	i, o = []int{1, 2, 3, 2, 1, 2}, []int{1, 2, 3, 2, 2, 1}
	input_str = S(i)
	nextPermutation(i)
	T(_t, input_str, S(i), S(o))

	i, o = []int{1, 2, 3, 1}, []int{1, 3, 1, 2}
	input_str = S(i)
	nextPermutation(i)
	T(_t, input_str, S(i), S(o))

	i, o = []int{2, 3, 1, 3, 3}, []int{2, 3, 3, 1, 3}
	input_str = S(i)
	nextPermutation(i)
	T(_t, input_str, S(i), S(o))
}
