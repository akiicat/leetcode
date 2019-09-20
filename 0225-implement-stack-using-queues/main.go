package main
import "fmt"
import "strconv"

func main() {
  i1 := []string{"MyStack","push","push","top","pop","empty"}
  i2 := [][]int{[]int{},[]int{1},[]int{2},[]int{},[]int{},[]int{}}
  o := []string{"null","null","null","2","2","false"}

  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", run(i1, i2))
  fmt.Printf("Expect: %v\n", o)
}

func run(intrs []string, values [][]int) []string {
  var obj MyStack
  rtn := make([]string, len(intrs))

  for i, intr := range intrs {
    switch intr {
    case "MyStack":
      obj = Constructor()
      rtn[i] = "null"
    case "push":
      obj.Push(values[i][0])
      rtn[i] = "null"
    case "top":
      rtn[i] = strconv.Itoa(obj.Top())
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

