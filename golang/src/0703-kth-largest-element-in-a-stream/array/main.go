package main
import "sort"

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

