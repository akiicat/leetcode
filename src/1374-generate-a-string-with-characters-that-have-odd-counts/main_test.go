package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestGenerateTheString(_t *testing.T) {
	i, o := 4, "aaab"
	T(_t, S(i), S(valid(i, generateTheString(i), o)), S(o))

	i, o = 2, "ab"
	T(_t, S(i), S(valid(i, generateTheString(i), o)), S(o))

	i, o = 7, "aaaaaaa"
	T(_t, S(i), S(valid(i, generateTheString(i), o)), S(o))
}

func valid(n int, i, o string) string {
	if n != len(i) {
		return i
	}

	m := map[rune]int{}

	for _, v := range i {
		m[v]++
	}

	for _, v := range m {
		if v == 0 {
			continue
		}

		if v%2 != 1 {
			return i
		}
	}

	return o
}
