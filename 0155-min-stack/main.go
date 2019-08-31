package main
import "fmt"
import "strconv"

func main() {
  i1 := []string{"MinStack","push","push","push","getMin","pop","top","getMin"}
  i2 := [][]int{[]int{},[]int{-2},[]int{0},[]int{-3},[]int{},[]int{},[]int{},[]int{}}
  o := []string{"null","null","null","null","-3","null","0","-2"}

  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", run(i1, i2))
  fmt.Printf("Expect: %v\n", o)


  i1 = []string{"MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"}
  i2 = [][]int{[]int{},[]int{2147483646},[]int{2147483646},[]int{2147483647},[]int{},[]int{},[]int{},[]int{},[]int{},[]int{},[]int{2147483647},[]int{},[]int{},[]int{-2147483648},[]int{},[]int{},[]int{},[]int{}}

  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", run(i1, i2))
}

func run(intrs []string, values [][]int) []string {
  var obj MinStack
  rtn := make([]string, len(intrs))

  for i, intr := range intrs {
    switch intr {
    case "MinStack":
      obj = Constructor()
      rtn[i] = "null"
    case "push":
      obj.Push(values[i][0])
      rtn[i] = "null"
    case "getMin":
      rtn[i] = strconv.Itoa(obj.GetMin())
    case "pop":
      obj.Pop()
      rtn[i] = "null"
    case "top":
      rtn[i] = strconv.Itoa(obj.Top())
    default:
      fmt.Printf("Command Not Found\n")
      rtn[i] = "null"
    }
  }
  return rtn
}

// T: O(n)
// M: O(1)
// -- start --

type MinStack struct {
  arr []int
  min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
  return MinStack{}
}


func (this *MinStack) Push(x int)  {
  this.arr = append(this.arr, x)

  if len(this.min) == 0 || x <= this.min[len(this.min)-1] {
    this.min = append(this.min, x)
  }
}


func (this *MinStack) Pop()  {
  if len(this.arr) == 0 {
    return
  }

  if this.arr[len(this.arr)-1] == this.min[len(this.min)-1] {
    this.min = this.min[:len(this.min)-1]
  }

  this.arr = this.arr[:len(this.arr)-1]
}


func (this *MinStack) Top() int {
  if len(this.arr) == 0 {
    return 0
  }

  return this.arr[len(this.arr)-1]
}


func (this *MinStack) GetMin() int {
  if len(this.arr) == 0 {
    return 0
  }

  return this.min[len(this.min)-1]
}

// -- end --

