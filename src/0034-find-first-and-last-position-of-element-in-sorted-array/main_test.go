package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestSearchRange(_t *testing.T) {
	i, t, o := []int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}
	T(_t, S(i, t), S(searchRange(i, t)), S(o))

	i, t, o = []int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}
	T(_t, S(i, t), S(searchRange(i, t)), S(o))

	i, t, o = []int{5, 7, 7, 8, 8, 10}, 5, []int{0, 0}
	T(_t, S(i, t), S(searchRange(i, t)), S(o))

	i, t, o = []int{1}, 1, []int{0, 0}
	T(_t, S(i, t), S(searchRange(i, t)), S(o))

	i, t, o = []int{}, 1, []int{-1, -1}
	T(_t, S(i, t), S(searchRange(i, t)), S(o))
}
