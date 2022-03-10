package main
import "testing"
import . "main/pkg/testing_helper"
import "fmt"
import "strconv"

func TestMyStack(_t *testing.T) {
  i1 := []string{"MyStack","push","push","top","pop","empty"}
  i2 := [][]int{[]int{},[]int{1},[]int{2},[]int{},[]int{},[]int{}}
  o := []string{"null","null","null","2","2","false"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))
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
