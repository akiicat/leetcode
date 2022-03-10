package main
import "testing"
import . "main/pkg/testing_helper"
import "fmt"
import "strconv"

func TestMinStack(_t *testing.T) {
  i1 := []string{"MinStack","push","push","push","getMin","pop","top","getMin"}
  i2 := [][]int{[]int{},[]int{-2},[]int{0},[]int{-3},[]int{},[]int{},[]int{},[]int{}}
  o := []string{"null","null","null","null","-3","null","0","-2"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))

  i1 = []string{"MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"}
  i2 = [][]int{[]int{},[]int{2147483646},[]int{2147483646},[]int{2147483647},[]int{},[]int{},[]int{},[]int{},[]int{},[]int{},[]int{2147483647},[]int{},[]int{},[]int{-2147483648},[]int{},[]int{},[]int{},[]int{}}
  o = []string{"null", "null", "null", "null", "2147483647", "null", "2147483646", "null", "2147483646", "null", "null", "2147483647", "2147483647", "null", "-2147483648", "-2147483648", "null", "2147483647"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))
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
