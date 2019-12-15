package main
import "fmt"

func main() {
  i, o := []int{1,2,2,4}, []int{2,3}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", findErrorNums(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = []int{4,2,2,1}, []int{2,3}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", findErrorNums(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = []int{2,2}, []int{2,1}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", findErrorNums(i))
  fmt.Printf("Expect: %v\n", o)
}

// https://leetcode.com/articles/set-mismatch/
// T: O(n)
// M: O(n)
// -- start --

// XOR
// T: O(n)
// M: O(1)
func findErrorNums(nums []int) []int {
  xor, xor0, xor1 := 0, 0, 0

  for i, v := range nums {
    xor ^= (i+1) ^ v
  }

  rightmostbit := xor & ^(xor - 1)

  for i, v := range nums {
    if v & rightmostbit != 0 {
      xor1 ^= v
    } else {
      xor0 ^= v
    }

    if (i + 1) & rightmostbit != 0 {
      xor1 ^= (i + 1)
    } else {
      xor0 ^= (i + 1)
    }
  }

  for _, v := range nums {
    if v == xor0 {
      return []int{xor0, xor1}
    }
  }
  return []int{xor1, xor0}
}

// T: O(n)
// M: O(1)
func findErrorNumsMath(nums []int) []int {
    n := len(nums)                      // x for missed number, and y for twiced number
    a := n * (n+1) / 2                  // a = 1 + 2 + ... + n
    b := n * (n+1) * (2*n+1) / 6        // b = 1*1 + 2*2 + ... + n*n
    for _, i := range nums{
        a -= i
        b -= i * i
    }                                   // a = x - y, b = x*x - y*y
    b /= a                              // b = x + y
    a = (a+b) / 2                       // a = y
    b = b - a                           // b = x
    return []int{b, a}
}

// T: O(n)
// M: O(1)
func findErrorNumsNegative(nums []int) []int {
  dup, missing := -1, -1
  for _, v := range nums {
    if nums[Abs(v)-1] < 0 {
      dup = Abs(v)
    } else {
      nums[Abs(v)-1] *= -1
    }
  }

  for i := 0; i < len(nums); i++ {
    if nums[i] > 0 {
      missing = i + 1
    }
  }

  return []int{dup, missing}

}

func Abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

// T: O(n)
// M: O(n)
func findErrorNumsArray(nums []int) []int {
  arr := make([]bool, len(nums)+1)
  rtn := make([]int, 2)

  for _, v := range nums {
    if arr[v] {
      rtn[0] = v
    } else {
      arr[v] = true
    }
  }

  for i := 1; i < len(arr); i++ {
    if !arr[i] {
      rtn[1] = i
      break
    }
  }

  return rtn
}


// -- end --

