package main

// T: O(n ** 2) n for the number of points
// M: O(n)
// -- start --

// T: O(n ** 2)
// M: O(n)
func numberOfBoomerangs(points [][]int) int {
  n := len(points)

  sum := 0
  for i := 0; i < n; i++ {
    m := make(map[int]int)
    for j := 0; j < n; j++ {
      if i == j {
        continue
      }

      l := Len(points[i], points[j])

      // m[l] will not over 4, 0 <= m[l] <= 4
      if v, ok := m[l]; ok {
        sum += v * 2
        m[l]++
      } else {
        m[l] = 1
      }
    }
  }

  return sum
}

// T: O(n ** 2)
// M: O(n)
func numberOfBoomerangsNested(points [][]int) int {
  n := len(points)

  sum := 0
  for i := 0; i < n; i++ {
    m := make(map[int]int)

    for j := 0; j < n; j++ {
      m[Len(points[i], points[j])]++
    }

    for k, v := range m {
      if k != 0 && v >= 2 {
        sum += v * (v - 1)
      }
    }
  }

  return sum
}

func Len(a, b []int) int {
  x := b[0] - a[0]
  y := b[1] - a[1]
  return x * x + y * y
}

// -- end --

