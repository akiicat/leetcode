package main
import "fmt"
import . "tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("3,9,20,null,null,15,7"), 3
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %d\n", maxDepth(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(h)
// -- start --

func maxDepth(root *TreeNode) int {
  if root == nil {
    return 0
  }

  return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}


// -- end --

