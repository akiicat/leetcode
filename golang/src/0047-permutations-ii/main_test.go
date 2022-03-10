package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestPermuteUnique(_t *testing.T) {
	i, o := []int{1, 1, 2}, [][]int{
		[]int{1, 1, 2},
		[]int{1, 2, 1},
		[]int{2, 1, 1},
	}
	T(_t, S(i), S(Sort(permuteUnique(i))), S(Sort(o)))

	i, o = []int{1, 2, 3}, [][]int{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{2, 1, 3},
		[]int{2, 3, 1},
		[]int{3, 1, 2},
		[]int{3, 2, 1},
	}
	T(_t, S(i), S(Sort(permuteUnique(i))), S(Sort(o)))

	i, o = []int{2, 2, 1, 1}, [][]int{
		[]int{1, 1, 2, 2},
		[]int{1, 2, 1, 2},
		[]int{1, 2, 2, 1},
		[]int{2, 1, 1, 2},
		[]int{2, 1, 2, 1},
		[]int{2, 2, 1, 1},
	}
	T(_t, S(i), S(Sort(permuteUnique(i))), S(Sort(o)))
}
