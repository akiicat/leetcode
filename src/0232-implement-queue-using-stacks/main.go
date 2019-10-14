package main
import "fmt"
import "strconv"

func main() {
  i1 := []string{"MyQueue","push","push","peek","pop","empty"}
  i2 := [][]int{[]int{},[]int{1},[]int{2},[]int{},[]int{},[]int{}}
  o := []string{"null","null","null","1","1","false"}

  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", run(i1, i2))
  fmt.Printf("Expect: %v\n", o)
}

func run(intrs []string, values [][]int) []string {
  var obj MyQueue
  rtn := make([]string, len(intrs))

  for i, intr := range intrs {
    switch intr {
    case "MyQueue":
      obj = Constructor()
      rtn[i] = "null"
    case "push":
      obj.Push(values[i][0])
      rtn[i] = "null"
    case "peek":
      rtn[i] = strconv.Itoa(obj.Peek())
    case "pop":
      rtn[i] = strconv.Itoa(obj.Pop())
    case "empty":
      rtn[i] = strconv.FormatBool(obj.Empty())
    default:
      fmt.Printf("Command Not Found\n")
      rtn[i] = "null"
    }
  }
  return rtn
}

// leetcode 155. 225.
// T: O(n)
// M: O(1)
// -- start --

type MyQueue struct {
  arr []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
  return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
  this.arr = append(this.arr, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
  rtn := this.arr[0]
  this.arr = this.arr[1:]
  return rtn
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
  return this.arr[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
  return len(this.arr) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

// -- end --

