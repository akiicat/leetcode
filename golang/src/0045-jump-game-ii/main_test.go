package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestJump(_t *testing.T) {
	i, o := []int{2, 3, 1, 1, 4}, 2
	T(_t, S(i), S(jump(i)), S(o))

	i, o = []int{2, 3, 0, 1, 4}, 2
	T(_t, S(i), S(jump(i)), S(o))

	i, o = []int{1, 2, 1, 1, 1}, 3
	T(_t, S(i), S(jump(i)), S(o))

	i, o = []int{2, 1}, 1
	T(_t, S(i), S(jump(i)), S(o))

	i, o = []int{0}, 0
	T(_t, S(i), S(jump(i)), S(o))
}

func TestJumpDP(_t *testing.T) {
	i, o := []int{2, 3, 1, 1, 4}, 2
	T(_t, S(i), S(jumpDP(i)), S(o))

	i, o = []int{2, 3, 0, 1, 4}, 2
	T(_t, S(i), S(jumpDP(i)), S(o))

	i, o = []int{1, 2, 1, 1, 1}, 3
	T(_t, S(i), S(jumpDP(i)), S(o))

	i, o = []int{2, 1}, 1
	T(_t, S(i), S(jump(i)), S(o))

	i, o = []int{0}, 0
	T(_t, S(i), S(jump(i)), S(o))
}
