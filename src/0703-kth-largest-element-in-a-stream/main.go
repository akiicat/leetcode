package main
import "fmt"
import "sort"
import "strconv"

func main() {
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

  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", run(i1, i2))
  fmt.Printf("Expect: %v\n", o)

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

  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", run(i1, i2))
  fmt.Printf("Expect: %v\n", o)

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

  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", run(i1, i2))
  fmt.Printf("Expect: %v\n", o)
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

// T: O(n * log(k))
// M: O(k)
// -- start --

type KthLargest struct {
  k int
  arr []int
}


func Constructor(k int, nums []int) KthLargest {
  sort.Sort(sort.Reverse(sort.IntSlice(nums)))
  return KthLargest{k, nums}
}


func (this *KthLargest) Add(val int) int {
  this.arr = append(this.arr, val)

  for i, v := range this.arr {
    if v >= val {
      continue
    }

    copy(this.arr[i+1:], this.arr[i:])
    this.arr[i] = val

    break
  }

  if len(this.arr) < this.k {
    return -1
  }

  this.arr = this.arr[:this.k]

  return this.arr[this.k-1]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// -- end --

