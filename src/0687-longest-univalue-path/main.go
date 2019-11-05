package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("5,4,5,1,1,5"), 2
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %d\n", longestUnivaluePath(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = NewTreeNode("1,4,5,4,4,5"), 2
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %d\n", longestUnivaluePath(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(h)
// -- start --

func longestUnivaluePath(root *TreeNode) int {
  if root == nil {
    return 0
  }

  max := 0
  BFS(root, root.Val, &max)
  return max
}

func BFS(root *TreeNode, prev int, max *int) int {
  if root == nil {
    return 0
  }

  l := BFS(root.Left, root.Val, max)
  r := BFS(root.Right, root.Val, max)
  *max = Max(*max, l + r)

  if root.Val != prev {
    return 0
  }

  return Max(l, r) + 1
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

