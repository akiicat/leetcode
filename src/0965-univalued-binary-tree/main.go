package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// leetcode 572.
// T: O(n)
// M: O(h)
// -- start --

func isUnivalTree(root *TreeNode) bool {
  if root == nil {
    return true
  }

  if root.Left != nil && root.Val != root.Left.Val {
    return false
  }

  if root.Right != nil && root.Val != root.Right.Val {
    return false
  }

  return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}

// -- end --

