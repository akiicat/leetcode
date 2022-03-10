package main
import "testing"
import . "main/pkg/testing_helper"
import "fmt"
import "strconv"

func TestNumArray(_t *testing.T) {
  i1 := []string{"NumArray","sumRange","sumRange","sumRange"}
  i2 := [][]int{[]int{-2,0,3,-5,2,-1},[]int{0,2},[]int{2,5},[]int{0,5}}
  o := []string{"null","1","-1","-3"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))
}

func run(intrs []string, values [][]int) []string {
  var obj NumArray
  rtn := make([]string, len(intrs))

  for i, intr := range intrs {
    switch intr {
    case "NumArray":
      obj = Constructor(values[i])
      rtn[i] = "null"
    case "sumRange":
      rtn[i] = strconv.Itoa(obj.SumRange(values[i][0], values[i][1]))
    default:
      fmt.Printf("Command Not Found\n")
      rtn[i] = "null"
    }
  }
  return rtn
}
