package main

import (
	. "main/pkg/testing_helper"
	"testing"
)

func TestGroupAnagrams(_t *testing.T) {
	i, o := []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[][]string{
			[]string{"ate", "eat", "tea"},
			[]string{"nat", "tan"},
			[]string{"bat"},
		}
	T(_t, S(i), S(Sort(groupAnagrams(i))), S(Sort(o)))
}
