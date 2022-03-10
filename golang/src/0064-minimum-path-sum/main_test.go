package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestMinPathSum(_t *testing.T) {
	i, o := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}, 7
	T(_t, S(i), S(minPathSum(i)), S(o))

	i, o = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}, 12
	T(_t, S(i), S(minPathSum(i)), S(o))
}
