package main

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

