package main

// T: O(n) 
// M: O(1)
// -- start --

func canPlaceFlowers(flowerbed []int, n int) bool {
  count, mask := 0, 0b010

  for _, v := range flowerbed {
    mask = (mask & 0b011 << 1) | v
    if mask == 0 {
      count++
      mask = 0b010
    }
  }

  if mask & 0b011 == 0 {
    count++
  }

  return count >= n
}

// -- end --

