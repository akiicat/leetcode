package main
import "testing"
import . "main/pkg/testing_helper"

func TestReorderLogFiles(_t *testing.T) {
  i, o := []string{"dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"}, []string{"let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"}
  T(_t, S(i), S(reorderLogFiles(i)), S(o))

  i, o = []string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo","a2 act car"}, []string{"a2 act car","g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"}
  T(_t, S(i), S(reorderLogFiles(i)), S(o))
}

