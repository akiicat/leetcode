package main
import "testing"
import . "main/pkg/testing_helper"

func TestNumSmallerByFrequency(t *testing.T) {
  queries, words, o := []string{"cbd"}, []string{"zaaaz"}, []int{1}
  T(t, S(queries, words), S(numSmallerByFrequency(queries, words)), S(o))

  queries, words, o = []string{"bbb","cc"}, []string{"a","aa","aaa","aaaa"}, []int{1,2}
  T(t, S(queries, words), S(numSmallerByFrequency(queries, words)), S(o))
}

