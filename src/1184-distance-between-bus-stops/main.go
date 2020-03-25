package main

// T: O(n)
// M: O(1)
// -- start --

func distanceBetweenBusStops(distance []int, start int, destination int) int {
  sum := 0
  for _, v := range distance {
    sum += v
  }

  s := 0
  for start != destination {
    s += distance[start]
    start = (start + 1) % len(distance)
  }
  return Min(s, sum-s)
}

func Min(a, b int) int {
  if a > b {
    return b
  }
  return a
}

// -- end --

