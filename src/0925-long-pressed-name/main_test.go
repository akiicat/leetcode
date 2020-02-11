package main
import "testing"
import . "main/pkg/testing_helper"

func TestIsLongPressedName(t *testing.T) {
  name, typed, o := "alex", "aaleex", true
  T(t, S(name, typed), S(isLongPressedName(name, typed)), S(o))

  name, typed, o = "saeed", "ssaaedd", false
  T(t, S(name, typed), S(isLongPressedName(name, typed)), S(o))

  name, typed, o = "leelee", "lleeelee", true
  T(t, S(name, typed), S(isLongPressedName(name, typed)), S(o))

  name, typed, o = "vtkgn", "vttkgnn", true
  T(t, S(name, typed), S(isLongPressedName(name, typed)), S(o))
}

