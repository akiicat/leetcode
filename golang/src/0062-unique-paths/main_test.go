package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestUniquePaths(_t *testing.T) {
	m, n, o := 3, 7, 28
	T(_t, S(m, n), S(uniquePaths(m, n)), S(o))

	m, n, o = 3, 2, 3
	T(_t, S(m, n), S(uniquePaths(m, n)), S(o))

	m, n, o = 7, 3, 28
	T(_t, S(m, n), S(uniquePaths(m, n)), S(o))

	m, n, o = 3, 3, 6
	T(_t, S(m, n), S(uniquePaths(m, n)), S(o))
}
