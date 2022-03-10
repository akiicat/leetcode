package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// T: O(n)
// M: O(h)
// -- start --

func minDepth(root *TreeNode) int {
  if root == nil {
    return 0
  }

  if root.Left == nil && root.Right == nil {
    return 1
  }

  if root.Left == nil {
    return minDepth(root.Right) + 1
  }

  if root.Right == nil {
    return minDepth(root.Left) + 1
  }

  return Min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

