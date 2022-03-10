package main

// leetcode 155. 232.
// T: O(n)
// M: O(1)
// -- start --

type MyStack struct {
  arr []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
  return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
  this.arr = append(this.arr, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
  rtn := this.arr[len(this.arr)-1]
  this.arr = this.arr[:len(this.arr)-1]
  return rtn
}


/** Get the top element. */
func (this *MyStack) Top() int {
  return this.arr[len(this.arr)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
  return len(this.arr) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

// -- end --

