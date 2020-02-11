package main
import "testing"
import . "main/pkg/testing_helper"

func TestCheckRecord(_t *testing.T) {
  i, o := "PPALLP", true
  T(_t, S(i), S(checkRecord(i)), S(o))

  i, o = "PPALLL", false
  T(_t, S(i), S(checkRecord(i)), S(o))

  i, o = "LALL", true
  T(_t, S(i), S(checkRecord(i)), S(o))
}

func TestCheckRecordStrings(_t *testing.T) {
  i, o := "PPALLP", true
  T(_t, S(i), S(checkRecordStrings(i)), S(o))

  i, o = "PPALLL", false
  T(_t, S(i), S(checkRecordStrings(i)), S(o))

  i, o = "LALL", true
  T(_t, S(i), S(checkRecordStrings(i)), S(o))
}
