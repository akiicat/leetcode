package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestCanJump(_t *testing.T) {
	i, o := []int{2, 3, 1, 1, 4}, true
	T(_t, S(i), S(canJump(i)), S(o))

	i, o = []int{3, 2, 1, 0, 4}, false
	T(_t, S(i), S(canJump(i)), S(o))

}

func TestCanJumpDynamicProgramming(_t *testing.T) {
	i, o := []int{2, 3, 1, 1, 4}, true
	T(_t, S(i), S(canJumpDynamicProgramming(i)), S(o))

	i, o = []int{3, 2, 1, 0, 4}, false
	T(_t, S(i), S(canJumpDynamicProgramming(i)), S(o))
}
