package main

// T: O(n)
// M: O(n)
// -- start --

func numPrimeArrangements(n int) int {
  d := countPrimes(n+1)

  sum := 1
  for i := 1; i <= d; i++ {
    sum = (sum * i) % 1000000007
  }

  for i := 1; i <= n - d; i++ {
    sum = (sum * i) % 1000000007
  }

  return sum
}

// leetcode 204. count primes
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	isComposite := make([]bool, n)
	count := n / 2
	for i := 3; i*i < n; i += 2 {

		if isComposite[i] {
			continue
		}
		for j := i * i; j < n; j += 2 * i {
			if !isComposite[j] {
				count--
				isComposite[j] = true
			}
		}
	}
	return count
}

// -- end --

