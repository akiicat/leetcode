package main
import "fmt"
// import "math/bits"

func main() {
  fmt.Printf("Input:  3, 2, 3\nOutput: %d\nExpect: 3\n", majorityElement([]int{3, 2, 3}))
  fmt.Printf("Input:  2, 2, 1, 1, 1, 2, 2\nOutput: %d\nExpect: 2\n", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

// T: O(N)
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

