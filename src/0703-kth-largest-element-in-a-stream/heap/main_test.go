package main
import "testing"
import . "main/pkg/testing_helper"
import "fmt"
import "strconv"

func TestKthLargest(_t *testing.T) {
  i1 := []string{"KthLargest","add","add","add","add","add"}
  i2 := []interface{}{
    []interface{}{3,[]int{4,5,8,2}},
    []int{3},
    []int{5},
    []int{10},
    []int{9},
    []int{4},
  }
  o := []string{"null","4","5","5","8","8"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))

  i1 = []string{"KthLargest","add","add","add","add","add"}
  i2 = []interface{}{
    []interface{}{3,[]int{4}},
    []int{3},
    []int{5},
    []int{10},
    []int{9},
    []int{4},
  }
  o = []string{"null","-1","3","4","5","5"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))

  i1 = []string{"KthLargest","add","add","add","add","add"}
  i2 = []interface{}{
    []interface{}{1,[]int{}},
    []int{-3},
    []int{-2},
    []int{-4},
    []int{0},
    []int{4},
  }
  o = []string{"null","-3","-2","-2","0","4"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))
}

func run(intrs []string, values []interface{}) []string {
  var obj KthLargest
  rtn := make([]string, len(intrs))

  for i, intr := range intrs {
    switch intr {
    case "KthLargest":
      v := values[i].([]interface{})
      obj = Constructor(v[0].(int), v[1].([]int))
      rtn[i] = "null"
    case "add":
      v := values[i].([]int)
      rtn[i] = strconv.Itoa(obj.Add(v[0]))
    default:
      fmt.Printf("Command Not Found\n")
      rtn[i] = "null"
    }
  }
  return rtn
}

