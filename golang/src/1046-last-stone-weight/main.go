package main
import "sort"

// T: O(n)
// M: O(l) for 1000
// -- start --

func lastStoneWeight(stones []int) int {
  m := [1001]int{}
  for _, v := range stones {
    m[v]++
  }

  j := 0
  for i := len(m)-1; i > 0; i-- {
    if m[i] == 0 {
      continue
    }

    if j == 0 {
      if m[i] & 0x1 == 1 {
        j = i
      }
      m[i] = 0
      continue
    }

    m[i]--
    m[j-i]++

    i, j = Max(i, j-i) + 1, 0
  }

  return j
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// T: O(n * log(n))
// M: O(n)
func lastStoneWeightSort(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		diff := stones[len(stones)-1] - stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		stones = append(stones, diff)
	}
	return stones[0]
}

// -- end --

