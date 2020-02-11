package main

// T: O(n)
// M: O(1)
// -- start --

func canConstruct(ransomNote string, magazine string) bool {
  if len(ransomNote) > len(magazine) {
    return false
  }

  count := [26]int{}

  for _, c := range magazine {
    count[c - 'a']++;
  }

  for _, c := range ransomNote {
    count[c - 'a']--;
  }

  for _, num := range count {
    if num < 0 {
      return false
    }
  }

  return true
}

// -- end --

