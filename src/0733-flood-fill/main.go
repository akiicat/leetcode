package main
import "fmt"

func main() {
  image, sr, sc, newColor, o := [][]int{[]int{1,1,1}, []int{1,1,0}, []int{1,0,1}}, 1, 1, 2, [][]int{[]int{2,2,2}, []int{2,2,0}, []int{2,0,1}}
  fmt.Printf("Input:  %v, %d, %d, %d\n", image, sr, sc, newColor)
  fmt.Printf("Output: %v\n", floodFill(image, sr, sc, newColor))
  fmt.Printf("Expect: %v\n", o)

  image, sr, sc, newColor, o = [][]int{[]int{0,0,0}, []int{0,1,1}}, 1, 1, 1, [][]int{[]int{0,0,0}, []int{0,1,1}}
  fmt.Printf("Input:  %v, %d, %d, %d\n", image, sr, sc, newColor)
  fmt.Printf("Output: %v\n", floodFill(image, sr, sc, newColor))
  fmt.Printf("Expect: %v\n", o)
}

// T: O(n) n is the number of pixels
// M: O(n)
// -- start --

// BFS
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
  oldColor := image[sr][sc]

  if oldColor == newColor {
    return image
  }

  stack := []int{sr, sc}

  lr, lc := len(image), len(image[0])

  for len(stack) != 0 {
    r, c := stack[0], stack[1]
    stack = stack[2:]

    image[r][c] = newColor

    // up
    if r != 0 && image[r-1][c] == oldColor {
      stack = append(stack, r-1, c)
    }

    // down
    if r != lr-1 && image[r+1][c] == oldColor {
      stack = append(stack, r+1, c)
    }

    // left
    if c != 0 && image[r][c-1] == oldColor {
      stack = append(stack, r, c-1)
    }

    // right
    if c != lc-1 && image[r][c+1] == oldColor {
      stack = append(stack, r, c+1)
    }
  }

  return image
}

// -- end --

