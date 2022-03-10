package main

// T: O(n)
// M: O(1)
// -- start --

func findTheDifference(s string, t string) byte {

  sum := byte(t[len(t)-1])

  for i := range s {
    sum += t[i] - s[i]
  }

  return sum
}

func findTheDifferenceMap(s string, t string) byte {
  m := make(map[byte]int)

  for i := 0; i < len(s); i++ {
    m[s[i]] += 1
  }

  for i := 0; i < len(t); i++ {
    m[t[i]] -= 1
  }

  for k, v := range m {
    if v != 0 {
      return k
    }
  }

  return 0
}

// -- end --
