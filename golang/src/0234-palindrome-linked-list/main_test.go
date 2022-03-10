package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestIsPalindrome(_t *testing.T) {
  i, o := NewListNode([]int{1,2}), false
  T(_t, S(i), S(isPalindrome(i)), S(o))

  i, o = NewListNode([]int{1,2,2,1}), true
  T(_t, S(i), S(isPalindrome(i)), S(o))

  i, o = NewListNode([]int{1,2,1}), true
  T(_t, S(i), S(isPalindrome(i)), S(o))

  i, o = NewListNode([]int{1,0,0}), false
  T(_t, S(i), S(isPalindrome(i)), S(o))
}

