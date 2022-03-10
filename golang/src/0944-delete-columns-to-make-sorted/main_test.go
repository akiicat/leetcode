package main
import "testing"
import . "main/pkg/testing_helper"

func TestMinDeletionSize(t *testing.T) {
  i, o := []string{"cba","daf","ghi"}, 1
  T(t, S(i), S(minDeletionSize(i)), S(o))

  i, o = []string{"a", "b"}, 0
  T(t, S(i), S(minDeletionSize(i)), S(o))

  i, o = []string{"zyx","wvu","tsr"}, 3
  T(t, S(i), S(minDeletionSize(i)), S(o))
}

