package main
import "strconv"
import "sort"

// leetcode 414. 628.
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

