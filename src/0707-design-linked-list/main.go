package main

// leetcode 155. 232.
// T: O(n)
// M: O(1)
// -- start --

type Node struct {
  Val int
  Next *Node
}

type MyLinkedList struct {
  head *Node
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
  return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
  head := this.head
  for index > 0 && head != nil {
    head = head.Next
    index--
  }

  if head == nil {
    return -1
  }

  return head.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
  head := this.head
  this.head = &Node{val, head}
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
  if this.head == nil {
    this.AddAtHead(val)
    return
  }

  head := this.head
  for head.Next != nil {
    head = head.Next
  }

  head.Next = &Node{val, nil}
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
  if index == 0 {
    this.head = &Node{val, this.head}
    return
  }

  head := this.head
  for index > 1 {
    head = head.Next
    index--
  }
  head.Next = &Node{val, head.Next}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
  if index == 0 {
    this.head = this.head.Next
    return
  }

  head := this.head
  for index > 1 && head != nil {
    head = head.Next
    index--
  }

  if head == nil || head.Next == nil {
    return
  }

  head.Next = head.Next.Next
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

// -- end --

