package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/list_node"

func TestHasCycle(_t *testing.T) {
  i, pos, o := NewListNode([]int{3,2,0,-4}), 1, true
  _i := S(i, pos)
  AddPos(i, pos)
  T(_t, _i, S(hasCycle(i)), S(o))

  i, pos, o = NewListNode([]int{1,2}), 0, true
  _i = S(i, pos)
  AddPos(i, pos)
  T(_t, _i, S(hasCycle(i)), S(o))

  i, pos, o = NewListNode([]int{1,2}), -1, false
  _i = S(i, pos)
  AddPos(i, pos)
  T(_t, _i, S(hasCycle(i)), S(o))
}

func TestHasCycleHash(_t *testing.T) {
  i, pos, o := NewListNode([]int{3,2,0,-4}), 1, true
  _i := S(i, pos)
  AddPos(i, pos)
  T(_t, _i, S(hasCycleHash(i)), S(o))

  i, pos, o = NewListNode([]int{1,2}), 0, true
  _i = S(i, pos)
  AddPos(i, pos)
  T(_t, _i, S(hasCycleHash(i)), S(o))

  i, pos, o = NewListNode([]int{1,2}), -1, false
  _i = S(i, pos)
  AddPos(i, pos)
  T(_t, _i, S(hasCycleHash(i)), S(o))
}

func AddPos(head *ListNode, pos int) {
  if pos < 0 {
    return
  }

  cur := head

  for {
    if pos == 0 {
      cur = head
    }

    if head.Next == nil {
      head.Next = cur
      return
    }

    head = head.Next
  }
}
