package main
import "fmt"
import "strconv"

func main() {
  i1 := []string{"NumArray","sumRange","sumRange","sumRange"}
  i2 := [][]int{[]int{-2,0,3,-5,2,-1},[]int{0,2},[]int{2,5},[]int{0,5}}
  o := []string{"null","1","-1","-3"}

  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", run(i1, i2))
  fmt.Printf("Expect: %v\n", o)
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

// T: O(n)
// M: O(1)
// -- start --

type NumArray struct {
  arr []int
}


func Constructor(nums []int) NumArray {
  obj := NumArray{}

  if len(nums) == 0 {
    return obj
  }

  obj.arr = make([]int, len(nums))
  obj.arr[0] = nums[0]

  for i := 1; i < len(nums); i++ {
    obj.arr[i] = obj.arr[i-1] + nums[i]
  }

  return obj
}


func (this *NumArray) SumRange(i int, j int) int {
  if i == 0 {
    return this.arr[j]
  }

  return this.arr[j] - this.arr[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

// -- end --

