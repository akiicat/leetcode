package main

// T: O(n)
// M: O(n)
// -- start --

// T: O(n)
// M: O(n)
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

// T: O(n ** 2)
// M: O(1)
func countPrimesIsPrime(n int) int {
  count := 0

  for i := 1; i < n; i++ {
    if isPrime(i) {
      count++
    }
  }

  return count;
}

func isPrime(num int) bool {
  if num <= 1 {
    return false
  }

  for i := 2; i * i <= num; i++ {
    if num % i == 0 {
      return false
    }
  }

  return true
}

// -- end --
