package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestSearch(_t *testing.T) {
	i, t, o := []int{4, 5, 6, 7, 0, 1, 2}, 0, 4
	T(_t, S(i, t), S(search(i, t)), S(o))

	i, t, o = []int{4, 5, 6, 7, 0, 1, 2}, 3, -1
	T(_t, S(i, t), S(search(i, t)), S(o))

	i, t, o = []int{4, 5, 6, 7, 0, 1, 2}, 1, 5
	T(_t, S(i, t), S(search(i, t)), S(o))

	i, t, o = []int{1, 3, 5, 7, 9}, 9, 4
	T(_t, S(i, t), S(search(i, t)), S(o))

	i, t, o = []int{1, 3, 5}, 5, 2
	T(_t, S(i, t), S(search(i, t)), S(o))

	i, t, o = []int{1, 3, 5}, 0, -1
	T(_t, S(i, t), S(search(i, t)), S(o))

	i, t, o = []int{3, 5, 1}, 1, 2
	T(_t, S(i, t), S(search(i, t)), S(o))

	i, t, o = []int{1, 3}, 3, 1
	T(_t, S(i, t), S(search(i, t)), S(o))

	i, t, o = []int{1, 3}, 2, -1
	T(_t, S(i, t), S(search(i, t)), S(o))

	i, t, o = []int{1}, 1, 0
	T(_t, S(i, t), S(search(i, t)), S(o))

	i, t, o = []int{}, 5, -1
	T(_t, S(i, t), S(search(i, t)), S(o))
}

func TestSearchLinear(_t *testing.T) {
	i, t, o := []int{4, 5, 6, 7, 0, 1, 2}, 0, 4
	T(_t, S(i, t), S(searchLinear(i, t)), S(o))

	i, t, o = []int{4, 5, 6, 7, 0, 1, 2}, 3, -1
	T(_t, S(i, t), S(searchLinear(i, t)), S(o))

	i, t, o = []int{4, 5, 6, 7, 0, 1, 2}, 1, 5
	T(_t, S(i, t), S(searchLinear(i, t)), S(o))

	i, t, o = []int{1, 3, 5, 7, 9}, 9, 4
	T(_t, S(i, t), S(searchLinear(i, t)), S(o))

	i, t, o = []int{1, 3, 5}, 5, 2
	T(_t, S(i, t), S(searchLinear(i, t)), S(o))

	i, t, o = []int{1, 3, 5}, 0, -1
	T(_t, S(i, t), S(searchLinear(i, t)), S(o))

	i, t, o = []int{3, 5, 1}, 1, 2
	T(_t, S(i, t), S(searchLinear(i, t)), S(o))

	i, t, o = []int{1, 3}, 3, 1
	T(_t, S(i, t), S(searchLinear(i, t)), S(o))

	i, t, o = []int{1, 3}, 2, -1
	T(_t, S(i, t), S(search(i, t)), S(o))

	i, t, o = []int{1}, 1, 0
	T(_t, S(i, t), S(searchLinear(i, t)), S(o))

	i, t, o = []int{}, 5, -1
	T(_t, S(i, t), S(searchLinear(i, t)), S(o))
}
