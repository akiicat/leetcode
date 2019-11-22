package main
import "fmt"
import "sort"

func main() {
  i, o := []int{1,4,3,2}, 4
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", arrayPairSum(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(r) r is the maximum number in array
// M: O(r)
// -- start --

// T: O(r)
// M: O(r)
func arrayPairSum(nums []int) int {
  m := make([]int, 20001)

  for _, num := range nums {
    m[num + 10000]++
  }

  odd := true
  sum := 0

  for i := 0; i < 20001; i++ {
    for m[i] != 0 {
      if odd {
        sum += i - 10000
      }
      m[i]--
      odd = !odd
    }
  }

  return sum
}


// T: O(n * log(n))
// M: O(1)
func arrayPairSumSort(nums []int) int {
  sort.Ints(nums)

  sum := 0
  for i := 0; i < len(nums); i += 2 {
    sum += nums[i]
  }

  return sum
}

// -- end --

