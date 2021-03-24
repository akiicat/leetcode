package main

import (
	. "main/pkg/list_node"
	. "main/pkg/testing_helper"
	"testing"
)

func TestRotateRight(t *testing.T) {
	i, k, o := NewListNode([]int{1, 2, 3, 4, 5}), 2, NewListNode([]int{4, 5, 1, 2, 3})
	T(t, S(i), S(rotateRight(i, k)), S(o))

	i, k, o = NewListNode([]int{0, 1, 2}), 4, NewListNode([]int{2, 0, 1})
	T(t, S(i), S(rotateRight(i, k)), S(o))

	i, k, o = NewListNode([]int{0, 1, 2}), 3, NewListNode([]int{0, 1, 2})
	T(t, S(i), S(rotateRight(i, k)), S(o))

	i, k, o = NewListNode([]int{}), 0, NewListNode([]int{})
	T(t, S(i), S(rotateRight(i, k)), S(o))
}
