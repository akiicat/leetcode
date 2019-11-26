package main
import "fmt"
import "sort"

func main() {
  i, o := []int{2,6,4,8,10,9,15}, 5
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", findUnsortedSubarray(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1}, 0
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", findUnsortedSubarray(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,3,2,2,2}, 4
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", findUnsortedSubarray(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,2,3,3,3}, 0
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", findUnsortedSubarray(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{1,3,2,3,3}, 2
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", findUnsortedSubarray(i))
  fmt.Printf("Expect: %d\n", o)
}

// https://leetcode.com/articles/shortest-unsorted-continous-subarray/
// T: O(n)
// M: O(1)
// -- start --

// T: O(n)
// M: O(1)
func findUnsortedSubarray(nums []int) int {
  min, max := 1<<31, -1<<31
  flag := false
  for i := 1; i < len(nums); i++ {
    if nums[i-1] > nums[i] {
      flag = true
    }
    if flag {
      min = Min(min, nums[i]);
    }
  }

  flag = false
  for i := len(nums) - 2; i >= 0; i-- {
    if nums[i] > nums[i+1] {
      flag = true
    }
    if flag {
      max = Max(max, nums[i])
    }
  }

  var l, r int
  for l = 0; l < len(nums); l++ {
    if nums[l] > min {
      break
    }
  }
  for r = len(nums)-1; r >= 0; r-- {
    if nums[r] < max {
      break
    }
  }

  if r < l {
    return 0
  }

  return r - l + 1
}

// T: O(n)
// M: O(n)
func findUnsortedSubarrayStack(nums []int) int {
  stack := []int{}

  t := 0
  l := len(nums)
  for i := 0; i < len(nums); i++ {
    for len(stack) != 0 && nums[stack[len(stack)-1]] > nums[i] {
      t, stack = stack[len(stack)-1], stack[:len(stack)-1]
      l = Min(l, t)
    }
    stack = append(stack, i)
  }

  stack = []int{}
  r := 0
  for i := len(nums)-1; i >= 0; i-- {
    for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] {
      t, stack = stack[len(stack)-1], stack[:len(stack)-1]
      r = Max(r, t)
    }
    stack = append(stack, i)
  }

  if r < l {
    return 0
  }

  return r - l + 1
}

// T: O(n * log(n))
// M: O(n)
func findUnsortedSubarraySort(nums []int) int {
  n := len(nums)

  if n < 2 {
    return 0
  }

  arr := make([]int, n)
  copy(arr, nums)
  sort.Ints(arr)

  l, r := n, 0
  for i, v := range nums {
    if v != arr[i] {
      r = Max(r, i)
      l = Min(l, i)
    }
  }

  if r < l {
    return 0
  }

  return r - l + 1
}

// T: O(n ^ 2)
// M: O(1)
func findUnsortedSubarrayBruteForce(nums []int) int {
  n := len(nums)

  if n < 2 {
    return 0
  }

  l, r := n, 0

  for i := 0; i < n - 1; i++ {
    for j := i + 1; j < n; j++ {
      if nums[j] < nums[i] {
        r = Max(r, j)
        l = Min(l, i)
      }
    }
  }

  if r < l {
    return 0
  }

  return r - l + 1
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

