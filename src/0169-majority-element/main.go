package main
import "fmt"

func main() {
  i, o := []int{3,2,3}, 3
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", majorityElement(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = []int{2,2,1,1,1,2,2}, 2
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %d\n", majorityElement(i))
  fmt.Printf("Expect: %d\n", o)
}

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

