package main

// T: O(log(n))
// M: O(1)
// -- start --

func subtractProductAndSum(n int) int {
  prod, sum := 1, 0

  for n != 0 {
    prod *= n % 10
    sum += n % 10
    n /= 10
  }

  return prod - sum
}

// -- end --

