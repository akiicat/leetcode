package main
import "testing"
import . "main/pkg/testing_helper"

func TestUncommonFromSentences(t *testing.T) {
  a, b, o := "this apple is sweet", "this apple is sour", []string{"sweet", "sour"}
  T(t, S(a, b), S(uncommonFromSentences(a, b)), S(o))

  a, b, o = "apple apple", "banana", []string{"banana"}
  T(t, S(a, b), S(uncommonFromSentences(a, b)), S(o))
}

