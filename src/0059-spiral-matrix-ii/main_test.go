package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestGenerateMatrix(_t *testing.T) {
	i, o := 1, [][]int{[]int{1}}
	T(_t, S(i), S(generateMatrix(i)), S(o))

	i, o = 3, [][]int{
		[]int{1, 2, 3},
		[]int{8, 9, 4},
		[]int{7, 6, 5},
	}
	T(_t, S(i), S(generateMatrix(i)), S(o))
}
