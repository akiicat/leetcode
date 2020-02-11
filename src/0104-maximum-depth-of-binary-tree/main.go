package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// leetcode 543.
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

