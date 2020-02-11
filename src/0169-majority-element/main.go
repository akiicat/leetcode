package main

// T: O(n)
// M: O(1)
// -- start --

func majorityElement(nums []int) int {
	var c, m int
	for _, v := range nums {
		if c == 0 {
			m = v
		}
		if v == m {
			c++
		} else {
			c--
		}
	}
	return m
}

func majorityElementMap(nums []int) int {
  m := make(map[int]int)

  for _, num := range(nums) {
    m[num]++
  }

  for k := range m {
    if m[k] > len(nums) / 2.0 {
      return k
    }
  }

  return 0
}


// -- end --

