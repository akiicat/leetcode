package main
import "fmt"

func main() {
  i, o := [][]int{[]int{1,1,1},[]int{1,0,1},[]int{1,1,1}}, [][]int{[]int{0,0,0},[]int{0,0,0},[]int{0,0,0}}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", imageSmoother(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = [][]int{[]int{1}}, [][]int{[]int{1}}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", imageSmoother(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = [][]int{[]int{2,3}}, [][]int{[]int{2,2}}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", imageSmoother(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = [][]int{[]int{2},[]int{3}}, [][]int{[]int{2},[]int{2}}
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %v\n", imageSmoother(i))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n) n is the number of pixels in image
// M: O(n)
// -- start --

func imageSmoother(M [][]int) [][]int {
  m, n := len(M), len(M[0])

  dx :=[]int{-1,-1,-1, 0, 0, 1, 1, 1}
  dy :=[]int{-1, 0, 1,-1, 1,-1, 0, 1}

  for y := 0; y < m; y++ {
    for x := 0; x < n; x++ {
      sum, cnt := M[y][x], 1
      for i := 0; i < 8; i++ {
        ny := y + dy[i]
        nx := x + dx[i]
        if nx >= 0 && nx < n && ny >= 0 && ny < m {
          sum += M[ny][nx] & 0xff
          cnt++
        }
      }
      M[y][x] |= (sum / cnt) << 8
    }
  }

  for y := 0; y < m; y++ {
    for x := 0; x < n; x++ {
      M[y][x] >>= 8
    }
  }

  return M
}

// -- end --

