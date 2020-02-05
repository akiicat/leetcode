package main
import "fmt"

func main() {
  c, n, o := 7, 4, []int{1,2,3,1}
  fmt.Printf("Input:  %d %d\n", c, n)
  fmt.Printf("Output: %v\n", distributeCandies(c, n))
  fmt.Printf("Expect: %v\n", o)

  c, n, o = 10, 3, []int{5,2,3}
  fmt.Printf("Input:  %d %d\n", c, n)
  fmt.Printf("Output: %v\n", distributeCandies(c, n))
  fmt.Printf("Expect: %v\n", o)

  c, n, o = 12, 3, []int{5,4,3}
  fmt.Printf("Input:  %d %d\n", c, n)
  fmt.Printf("Output: %v\n", distributeCandies(c, n))
  fmt.Printf("Expect: %v\n", o)

  c, n, o = 21, 3, []int{5,7,9}
  fmt.Printf("Input:  %d %d\n", c, n)
  fmt.Printf("Output: %v\n", distributeCandies(c, n))
  fmt.Printf("Expect: %v\n", o)

  c, n, o = 23, 3, []int{7,7,9}
  fmt.Printf("Input:  %d %d\n", c, n)
  fmt.Printf("Output: %v\n", distributeCandies(c, n))
  fmt.Printf("Expect: %v\n", o)

  c, n, o = 60, 4, []int{15,18,15,12}
  fmt.Printf("Input:  %d %d\n", c, n)
  fmt.Printf("Output: %v\n", distributeCandies(c, n))
  fmt.Printf("Expect: %v\n", o)

  c, n, o = 90, 4, []int{27,18,21,24}
  fmt.Printf("Input:  %d %d\n", c, n)
  fmt.Printf("Output: %v\n", distributeCandies(c, n))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(log(c))
// M: O(1)
// -- start --

// T: O(log(c))
// M: O(1)
func distributeCandies(candies int, num_people int) []int {
  m := make([]int, num_people)

  for i := 0; candies > 0; i++ {
    m[i % num_people] += Min(candies, i+1)
    candies -= i+1
  }

  return m
}

func Min(a, b int) int {
  if a > b {
    return b
  }
  return a
}

// T: O(n + log(c))
// M: O(1)
func distributeCandiesMath(candies int, num_people int) []int {
  m := make([]int, num_people)

  num := 0
  for candies > num {
    num++
    candies -= num

  }

  mod := num % num_people
  div := num / num_people

  for i := 0; i < num_people; i++ {
    if i < mod {
      m[i] += i + 1 + div * num_people
    }
    if i == mod {
      m[i] += candies
    }

    m[i] += div * (i + 1) + num_people * (((div-1) * div) / 2)
  }

  return m
}

// -- end --

