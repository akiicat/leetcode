package main
import "testing"
import . "main/pkg/testing_helper"

func TestCountCharacters(t *testing.T) {
  w, c, o := []string{"cat","bt","hat","tree"}, "atach", 6
  T(t, S(w, c), S(countCharacters(w, c)), S(o))
}

