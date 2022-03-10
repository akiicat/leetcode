package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestUniquePathsWithObstacles(_t *testing.T) {
	i, o := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}, 2
	T(_t, S(i), S(uniquePathsWithObstacles(i)), S(o))

	i, o = [][]int{
		[]int{0, 1},
		[]int{0, 0},
	}, 1
	T(_t, S(i), S(uniquePathsWithObstacles(i)), S(o))

	i, o = [][]int{
		[]int{1, 1},
	}, 0
	T(_t, S(i), S(uniquePathsWithObstacles(i)), S(o))
}
