package main
import "testing"
import . "main/pkg/testing_helper"

func TestReverseWords(_t *testing.T) {
  i, o := "Let's a ab take LeetCode contest", "s'teL a ba ekat edoCteeL tsetnoc"
  T(_t, S(i), S(reverseWords(i)), S(o))
}
