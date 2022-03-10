package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindWords(_t *testing.T) {
  i, o := []string{"Hello", "Alaska", "Dad", "Peace"}, []string{"Alaska", "Dad"}
  T(_t, S(i), S(findWords(i)), S(o))

  i, o = []string{"Aasdfghjkl","Qwertyuiop","zZxcvbnm"}, []string{"Aasdfghjkl","Qwertyuiop","zZxcvbnm"}
  T(_t, S(i), S(findWords(i)), S(o))
}
