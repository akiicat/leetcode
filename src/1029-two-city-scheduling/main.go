package main
import "sort"

// T: O(n)
// M: O(1)
// -- start --

// T: O(n * log(n))
// M: O(1)
func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})

  sum := 0
  for i := 0; i < len(costs); i++ {
		if i >= len(costs) / 2 {
			sum += costs[i][1]
			continue
		}
		sum += costs[i][0]
	}
  return sum
}

// T: O(n * log(n))
// M: O(n)
func twoCitySchedCostArray(costs [][]int) int {
  m := make([]int, len(costs))
  for i, v := range costs {
    m[i] = v[0] - v[1]
  }
  sort.Ints(m)

  s := 0
  for _, v := range costs {
    s += Min(v[0], v[1])
  }

  for _, v := range m[:len(costs)/2] {
    if v > 0 {
      s += v
    }
  }
  for _, v := range m[len(costs)/2:] {
    if v < 0 {
      s -= v
    }
  }

  return s
}

func Min(a, b int) int {
  if a > b {
    return b
  }
  return a
}

// -- end --

