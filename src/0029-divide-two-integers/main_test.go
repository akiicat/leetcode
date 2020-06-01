package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestRemoveDuplicates(_t *testing.T) {
	i1, i2, o := 10, 3, 3
	T(_t, S(i1, i2), S(divide(i1, i2)), S(o))

	i1, i2, o = -10, -3, 3
	T(_t, S(i1, i2), S(divide(i1, i2)), S(o))

	i1, i2, o = 7, -3, -2
	T(_t, S(i1, i2), S(divide(i1, i2)), S(o))

	i1, i2, o = 1, 1, 1
	T(_t, S(i1, i2), S(divide(i1, i2)), S(o))

	i1, i2, o = -2147483648, -1, 2147483647
	T(_t, S(i1, i2), S(divide(i1, i2)), S(o))

	i1, i2, o = -1010369383, -2147483648, 0
	T(_t, S(i1, i2), S(divide(i1, i2)), S(o))
}
