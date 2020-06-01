package main

import (
	. "main/pkg/list_node"
	. "main/pkg/testing_helper"
	"testing"
)

func TestSwapPairs(_t *testing.T) {
	i, o := NewListNode([]int{1, 2, 3, 4}), NewListNode([]int{2, 1, 4, 3})
	T(_t, S(i), S(swapPairs(i)), S(o))

	i, o = NewListNode([]int{1, 2, 3}), NewListNode([]int{2, 1, 3})
	T(_t, S(i), S(swapPairs(i)), S(o))

	i, o = NewListNode([]int{1, 2}), NewListNode([]int{2, 1})
	T(_t, S(i), S(swapPairs(i)), S(o))

	i, o = NewListNode([]int{1}), NewListNode([]int{1})
	T(_t, S(i), S(swapPairs(i)), S(o))

	i, o = NewListNode([]int{}), NewListNode([]int{})
	T(_t, S(i), S(swapPairs(i)), S(o))
}
