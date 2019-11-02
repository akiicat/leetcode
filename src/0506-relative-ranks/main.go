package main
import "fmt"
import "strconv"
import "sort"

func main() {
  i, o := []int{5,4,3,2,1}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", findRelativeRanks(i))
  fmt.Printf("Expect: %v\n", o)
}

// leetcode 414.
// T: O(n * log(n))
// M: O(n)
// -- start --

func findRelativeRanks(nums []int) []string {
  if len(nums) == 0 {
    return []string{}
  }

  m := make(map[int]int)

  for i, v := range nums {
    m[v] = i
  }

  sort.Ints(nums)

  rtn := make([]string, len(nums))
  rank := 1
  for i := len(nums) - 1; i >= 0; i-- {
    rtn[m[nums[i]]] = Rank(rank)
    rank++
  }

  return rtn
}

func Rank(rank int) string {
  switch rank {
  case 1:
    return "Gold Medal"
  case 2:
    return "Silver Medal"
  case 3:
    return "Bronze Medal"
  }

  return strconv.Itoa(rank)
}

// -- end --

