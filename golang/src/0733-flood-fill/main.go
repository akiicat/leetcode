package main

// T: O(n) n is the number of pixels
// M: O(n)
// -- start --

// BFS
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
  oldColor := image[sr][sc]

  if oldColor == newColor {
    return image
  }

  queue := []int{sr, sc}

  lr, lc := len(image), len(image[0])

  for len(queue) != 0 {
    r, c := queue[0], queue[1]
    queue = queue[2:]

    image[r][c] = newColor

    // up
    if r != 0 && image[r-1][c] == oldColor {
      queue = append(queue, r-1, c)
    }

    // down
    if r != lr-1 && image[r+1][c] == oldColor {
      queue = append(queue, r+1, c)
    }

    // left
    if c != 0 && image[r][c-1] == oldColor {
      queue = append(queue, r, c-1)
    }

    // right
    if c != lc-1 && image[r][c+1] == oldColor {
      queue = append(queue, r, c+1)
    }
  }

  return image
}

// -- end --

