package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/nested_integer"

func TestDeserialize(t *testing.T) {
  i, o := "[123,[456,[789]]]", NewNestedInteger("[123,[456,[789]]]")
  T(t, S(i), S(deserialize(i)), S(o))

  i, o = "324", NewNestedInteger("324")
  T(t, S(i), S(deserialize(i)), S(o))

  i, o = "[]", NewNestedInteger("[]")
  T(t, S(i), S(deserialize(i)), S(o))
}

