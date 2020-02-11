package main
import "testing"
import . "main/pkg/testing_helper"

func TestUniqueMorseRepresentations(_t *testing.T) {
  i, o := []string{"gin", "zen", "gig", "msg"}, 2
  T(_t, S(i), S(uniqueMorseRepresentations(i)), S(o))
}

