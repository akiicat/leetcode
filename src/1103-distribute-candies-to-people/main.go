package main

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

