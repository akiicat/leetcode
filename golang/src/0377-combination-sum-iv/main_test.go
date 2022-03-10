package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestCombinationSum4(_t *testing.T) {
	i, t, o := []int{1, 2, 3}, 4, 7
	T(_t, S(i, t), S(combinationSum4(i, t)), S(o))

	i, t, o = []int{2, 1, 3}, 35, 1132436852
	T(_t, S(i, t), S(combinationSum4(i, t)), S(o))

	i, t, o = []int{3, 33, 333}, 10000, 0
	T(_t, S(i, t), S(combinationSum4(i, t)), S(o))
}

func TestCombinationSum4Recursive(_t *testing.T) {
	i, t, o := []int{1, 2, 3}, 4, 7
	T(_t, S(i, t), S(combinationSum4Recursive(i, t)), S(o))

	i, t, o = []int{2, 1, 3}, 35, 1132436852
	T(_t, S(i, t), S(combinationSum4Recursive(i, t)), S(o))

	i, t, o = []int{3, 33, 333}, 10000, 0
	T(_t, S(i, t), S(combinationSum4Recursive(i, t)), S(o))
}
