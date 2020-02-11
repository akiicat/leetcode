package main

// T: O(n)
// M: O(1)
// -- start --

func lemonadeChange(bills []int) bool {
  five, ten := 0, 0

  for _, v := range bills {
    switch v {
    case 5:
      five++
    case 10:
      ten++
      five--
    case 20:
      if ten >= 1 {
        ten--
        five--
      } else {
        five -= 3
      }
    }

    if five < 0 || ten < 0 {
      return false
    }
  }

  return true
}

// -- end --

