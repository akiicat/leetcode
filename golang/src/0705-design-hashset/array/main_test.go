package main
import "testing"
import . "main/pkg/testing_helper"
import "fmt"

func TestMyHashSet(_t *testing.T) {
  i1 := []string{"MyHashSet","add","add","contains","contains","add","contains","remove","contains"}
  i2 := [][]int{[]int{},[]int{1},[]int{2},[]int{1},[]int{3},[]int{2},[]int{2},[]int{2},[]int{2}}
  o := []string{"null","null","null","true","false","null","true","null","false"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))
}

func run(intrs []string, values [][]int) []string {
  var obj MyHashSet
  rtn := make([]string, len(intrs))

  for i, intr := range intrs {
    switch intr {
    case "MyHashSet":
      obj = Constructor()
      rtn[i] = "null"
    case "add":
      obj.Add(values[i][0])
      rtn[i] = "null"
    case "contains":
      rtn[i] = fmt.Sprintf("%t", obj.Contains(values[i][0]))
    case "remove":
      obj.Remove(values[i][0])
      rtn[i] = "null"
    default:
      fmt.Printf("Command Not Found\n")
      rtn[i] = "null"
    }
  }
  return rtn
}

