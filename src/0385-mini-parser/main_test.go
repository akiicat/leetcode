package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/nested_integer"

func TestReverseBetween(t *testing.T) {
  i, o := "[123,[456,[789]]]", NewNestedInteger("[123,[456,[789]]]")
  T(t, S(i), S(deserialize(i)), S(o))

  i, o = "324", NewNestedInteger("324")
  T(t, S(i), S(deserialize(i)), S(o))
}

